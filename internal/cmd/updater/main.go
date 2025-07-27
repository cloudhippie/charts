package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cloudhippie/charts/internal/chart"
	"github.com/cloudhippie/charts/internal/chart/oci"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s <directory>\n", os.Args)
		os.Exit(2)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stdout,
			NoColor: false,
		},
	)

	if err := run(
		context.Background(),
		os.Args[1],
	); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context, root string) error {
	charts, err := chart.Find(root)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to find charts")

		return err
	}

	for _, c := range charts {
		if !c.Source.Disabled {
			log.Info().
				Str("chart", c.Meta.Name).
				Msg("checking version")

			latest, err := oci.Latest(
				path.Join(
					c.Source.Registry,
					c.Source.Image,
				),
				c.Source.Ignored,
				c.Source.Prerelease,
			)

			if err != nil {
				log.Error().
					Err(err).
					Str("chart", c.Meta.Name).
					Msg("failed to check latest version")

				continue
			}

			if c.Meta.AppVersion != latest.String() {
				log.Warn().
					Str("chart", c.Meta.Name).
					Str("old", c.Meta.AppVersion).
					Str("new", latest.String()).
					Msg("found newer version")

				if err := chart.UpdateApp(
					c,
					latest.String(),
				); err != nil {
					log.Error().
						Err(err).
						Str("chart", c.Meta.Name).
						Str("old", c.Meta.AppVersion).
						Str("new", latest.String()).
						Msg("failed to update version")
				}
			} else {
				log.Debug().
					Str("chart", c.Meta.Name).
					Str("version", c.Meta.AppVersion).
					Msg("no newer version found")
			}
		}

		for _, u := range c.Updates {
			log.Info().
				Str("chart", c.Meta.Name).
				Str("dependency", u.Name).
				Msg("checking version")

			latest, err := oci.Latest(
				u.Image,
				[]string{},
				false,
			)

			if err != nil {
				log.Error().
					Err(err).
					Str("chart", c.Meta.Name).
					Str("dependency", u.Name).
					Msg("failed to check latest version")

				continue
			}

			log.Debug().
				Str("chart", c.Meta.Name).
				Str("dependency", u.Name).
				Str("old", u.Tag).
				Str("new", latest.Original()).
				Msg("")

			if u.Tag != latest.Original() {
				log.Warn().
					Str("chart", c.Meta.Name).
					Str("dependency", u.Name).
					Str("old", u.Tag).
					Str("new", latest.Original()).
					Msg("found newer version")

				if err := chart.UpdateValue(
					c,
					u,
					latest.Original(),
				); err != nil {
					log.Error().
						Err(err).
						Str("chart", c.Meta.Name).
						Str("dependency", u.Name).
						Str("old", u.Tag).
						Str("new", latest.Original()).
						Msg("failed to update version")
				}
			} else {
				log.Debug().
					Str("chart", c.Meta.Name).
					Str("dependency", u.Name).
					Str("version", u.Tag).
					Msg("no newer version found")
			}
		}

		for _, d := range c.Meta.Dependencies {
			log.Info().
				Str("chart", c.Meta.Name).
				Str("dependency", d.Name).
				Msg("checking version")

			switch {
			case strings.HasPrefix(d.Repository, "oci://"):
				latest, err := oci.Latest(
					path.Join(
						strings.TrimPrefix(
							d.Repository,
							"oci://",
						),
						d.Name,
					),
					[]string{},
					false,
				)

				if err != nil {
					log.Error().
						Err(err).
						Str("chart", c.Meta.Name).
						Str("dependency", d.Name).
						Msg("failed to check latest version")

					continue
				}

				if d.Version != latest.Original() {
					log.Warn().
						Str("chart", c.Meta.Name).
						Str("dependency", d.Name).
						Str("old", d.Version).
						Str("new", latest.Original()).
						Msg("found newer version")

					if err := chart.UpdateDependency(
						c,
						d,
						latest.Original(),
					); err != nil {
						log.Error().
							Err(err).
							Str("chart", c.Meta.Name).
							Str("dependency", d.Name).
							Str("old", d.Version).
							Str("new", latest.Original()).
							Msg("failed to update version")
					}
				} else {
					log.Debug().
						Str("chart", c.Meta.Name).
						Str("dependency", d.Name).
						Str("version", d.Version).
						Msg("no newer version found")
				}
			default:
				log.Error().
					Str("chart", c.Meta.Name).
					Str("dependency", d.Name).
					Str("repo", d.Repository).
					Msg("unknown repository type")
			}
		}

		if err := chart.UpdateVersion(
			c,
		); err != nil {
			log.Error().
				Err(err).
				Str("chart", c.Meta.Name).
				Msg("failed to bump chart version")
		}

		if err := chart.CommitBump(
			c,
		); err != nil {
			log.Error().
				Err(err).
				Str("chart", c.Meta.Name).
				Msg("failed to commit chart version")
		}
	}

	return nil
}
