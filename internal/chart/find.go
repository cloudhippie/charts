package chart

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
)

var (
	ErrFilesNotFound = errors.New("chart files not found")
)

type Source struct {
	Registry   string   `yaml:"registry"`
	Image      string   `yaml:"image"`
	Ignored    []string `yaml:"ignored"`
	Disabled   bool     `yaml:"disabled"`
	Prerelease bool     `yaml:"prerelease"`
}

type Chart struct {
	Path    string
	Bump    Bump
	Changed bool
	Meta    *chart.Metadata
	Source  *Source
}

func Save(chart *Chart) error {
	if !chart.Changed {
		return nil
	}

	return chartutil.SaveChartfile(
		chart.Path,
		chart.Meta,
	)
}

func Find(root string) ([]*Chart, error) {
	files, err := gatherFiles(root)

	if err != nil {
		return nil, err
	}

	charts := make([]*Chart, 0, len(files))

	for _, file := range files {
		chart, err := chartutil.LoadChartfile(file)

		if err != nil {
			return nil, err
		}

		source, err := loadSource(file)

		if err != nil {
			return nil, err
		}

		charts = append(charts, &Chart{
			Path:   file,
			Meta:   chart,
			Source: source,
		})
	}

	return charts, nil
}

func gatherFiles(root string) ([]string, error) {
	var (
		files []string
	)

	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			switch d.Name() {
			case ".":
			case "internal":
				return filepath.SkipDir
			default:
				if strings.HasPrefix(d.Name(), ".") {
					return filepath.SkipDir
				}
			}
		} else if d.Name() == "Chart.yaml" {
			files = append(files, path)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}

func loadSource(file string) (*Source, error) {
	result := &Source{}

	content, err := os.ReadFile(
		path.Join(
			path.Dir(
				file,
			),
			"source.yaml",
		),
	)

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(
		content,
		result,
	); err != nil {
		return nil, err
	}

	return result, nil
}
