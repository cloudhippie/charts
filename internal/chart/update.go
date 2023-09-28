package chart

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/hashicorp/go-version"
	"github.com/rs/zerolog/log"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
)

type Bump int

const (
	NoBump Bump = iota
	PatchBump
	MinorBump
	MajorBump
)

func UpdateApp(c *Chart, latest string) error {
	old, err := version.NewVersion(
		c.Meta.AppVersion,
	)

	if err != nil {
		return err
	}

	new, err := version.NewVersion(
		latest,
	)

	if err != nil {
		return err
	}

	bump := detectBump(c, old, new)

	if c.Bump < bump {
		c.Bump = bump
		c.Changed = true
	}

	c.Meta.AppVersion = latest

	return nil
}

func UpdateDependency(c *Chart, d *chart.Dependency, latest string) error {
	old, err := version.NewVersion(
		d.Version,
	)

	if err != nil {
		return err
	}

	new, err := version.NewVersion(
		latest,
	)

	if err != nil {
		return err
	}

	bump := detectBump(c, old, new)

	if c.Bump < bump {
		c.Bump = bump
		c.Changed = true
	}

	d.Version = latest

	return nil
}

func UpdateVersion(c *Chart) error {
	if c.Bump == NoBump {
		return nil
	}

	if !c.Changed {
		return nil
	}

	old, err := version.NewVersion(
		c.Meta.Version,
	)

	if err != nil {
		return err
	}

	segments := old.Segments()

	switch {
	case c.Bump == MajorBump:
		segments[0] += 1
		segments[1] = 0
		segments[2] = 0
	case c.Bump == MinorBump:
		segments[1] += 1
		segments[2] = 0
	case c.Bump == PatchBump:
		segments[2] += 1
	}

	c.Meta.Version = fmt.Sprintf(
		"%d.%d.%d",
		segments[0],
		segments[1],
		segments[2],
	)

	if err := chartutil.SaveChartfile(
		c.Path,
		c.Meta,
	); err != nil {
		return err
	}

	if len(c.Meta.Dependencies) == 0 {
		return nil
	}

	helmPath, err := exec.LookPath("helm")

	if err != nil {
		return err
	}

	cmd := exec.Cmd{
		Path: helmPath,
		Args: []string{
			helmPath,
			"dependency",
			"update",
		},
		Dir:    path.Dir(c.Path),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Info().
		Str("chart", c.Meta.Name).
		Msg("executing helm dependency update")

	return cmd.Run()
}

func detectBump(c *Chart, old, new *version.Version) Bump {
	oldSegments := old.Segments()
	newSegments := new.Segments()

	if newSegments[0] > oldSegments[0] {
		log.Info().
			Str("chart", c.Meta.Name).
			Msg("detected major bump")

		return MajorBump
	}

	if newSegments[1] > oldSegments[1] {
		log.Info().
			Str("chart", c.Meta.Name).
			Msg("detected minor bump")

		return MinorBump
	}

	if newSegments[2] > oldSegments[2] {
		log.Info().
			Str("chart", c.Meta.Name).
			Msg("detected patch bump")

		return PatchBump
	}

	return NoBump
}
