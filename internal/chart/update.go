package chart

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/hashicorp/go-version"
	"github.com/rs/zerolog/log"
	chart "helm.sh/helm/v4/pkg/chart/v2"
	"helm.sh/helm/v4/pkg/chart/v2/util"
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

	if bump == NoBump {
		return nil
	}

	if err := util.SaveChartfile(
		c.Path,
		c.Meta,
	); err != nil {
		return err
	}

	switch bump {
	case MajorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"major(%s): updated application version %s",
				c.Meta.Name,
				c.Meta.AppVersion,
			),
		)
	case MinorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"minor(%s): updated application version %s",
				c.Meta.Name,
				c.Meta.AppVersion,
			),
		)
	case PatchBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"patch(%s): updated application version %s",
				c.Meta.Name,
				c.Meta.AppVersion,
			),
		)
	}

	return nil
}

func UpdateValue(c *Chart, d *Update, latest string) error {
	old, err := version.NewVersion(
		d.Tag,
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

	d.Tag = latest

	if bump == NoBump {
		return nil
	}

	c.Lines[d.Line] = updateMatch.ReplaceAllString(
		c.Lines[d.Line],
		fmt.Sprintf(
			"tag: %s  # updater: name=%s image=%s",
			d.Tag,
			d.Name,
			d.Image,
		),
	)

	if err := writeLines(
		path.Join(
			path.Dir(
				c.Path,
			),
			"values.yaml",
		),
		c.Lines,
	); err != nil {
		return err
	}

	switch bump {
	case MajorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"major(%s): updated %s tag to %s",
				c.Meta.Name,
				d.Name,
				d.Tag,
			),
		)
	case MinorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"minor(%s): updated %s tag to %s",
				c.Meta.Name,
				d.Name,
				d.Tag,
			),
		)
	case PatchBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"patch(%s): updated %s tag to %s",
				c.Meta.Name,
				d.Name,
				d.Tag,
			),
		)
	}

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

	if bump == NoBump {
		return nil
	}

	if err := util.SaveChartfile(
		c.Path,
		c.Meta,
	); err != nil {
		return err
	}

	if err := os.Remove(
		path.Join(
			path.Dir(
				c.Path,
			),
			"Chart.lock",
		),
	); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
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

	if err := cmd.Run(); err != nil {
		return err
	}

	switch bump {
	case MajorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"major(%s): updated dep %s to %s",
				c.Meta.Name,
				d.Name,
				d.Version,
			),
		)
	case MinorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"minor(%s): updated dep %s to %s",
				c.Meta.Name,
				d.Name,
				d.Version,
			),
		)
	case PatchBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"patch(%s): updated dep %s to %s",
				c.Meta.Name,
				d.Name,
				d.Version,
			),
		)
	}

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

	switch c.Bump {
	case MajorBump:
		segments[0] += 1
		segments[1] = 0
		segments[2] = 0
	case MinorBump:
		segments[1] += 1
		segments[2] = 0
	case PatchBump:
		segments[2] += 1
	}

	c.Meta.Version = fmt.Sprintf(
		"%d.%d.%d",
		segments[0],
		segments[1],
		segments[2],
	)

	return util.SaveChartfile(
		c.Path,
		c.Meta,
	)
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
