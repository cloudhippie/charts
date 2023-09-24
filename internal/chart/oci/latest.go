package oci

import (
	"errors"
	"regexp"
	"sort"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/hashicorp/go-version"
)

var (
	versionRegexp         *regexp.Regexp
	ErrNoVersionAvailable = errors.New("found not any version")
)

func init() {
	versionRegexp = regexp.MustCompile("^" + version.VersionRegexpRaw + "$")
}

func Latest(needle string, ignored []string, prerelease bool) (string, error) {
	repo, err := name.NewRepository(
		needle,
		name.WithDefaultRegistry(
			"docker.io",
		),
	)

	if err != nil {
		return "", err
	}

	tags, err := remote.List(
		repo,
	)

	if err != nil {
		return "", err
	}

	if len(tags) == 0 {
		return "", ErrNoVersionAvailable
	}

	versions := make([]*version.Version, 0)

	for _, tag := range tags {
		if ignoredTag(tag, ignored) {
			continue
		}

		if versionRegexp.FindStringSubmatch(tag) == nil {
			continue
		}

		v, err := version.NewVersion(tag)

		if err != nil {
			return "", err
		}

		if !prerelease && v.Prerelease() != "" {
			continue
		}

		versions = append(
			versions,
			v,
		)
	}

	sort.Sort(
		version.Collection(
			versions,
		),
	)

	return versions[len(versions)-1].String(), nil
}

func ignoredTag(tag string, ignored []string) bool {
	for _, i := range ignored {
		if tag == i {
			return true
		}
	}

	return false
}
