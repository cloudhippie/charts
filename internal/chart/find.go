package chart

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
)

var (
	ErrFilesNotFound = errors.New("chart files not found")
	updateMatch      = regexp.MustCompile(`tag:\s*(\S+)\s+#\s*updater:\s*name=([^\s]+)\s+image=([^\s]+)`)
)

type Source struct {
	Registry   string   `yaml:"registry"`
	Image      string   `yaml:"image"`
	Ignored    []string `yaml:"ignored"`
	Prerelease bool     `yaml:"prerelease"`
	Disabled   bool     `yaml:"disabled"`
}

type Update struct {
	Name  string
	Image string
	Tag   string
	Line  int
}

type Chart struct {
	Path    string
	Bump    Bump
	Changed bool
	Meta    *chart.Metadata
	Source  *Source
	Updates []*Update
	Lines   []string
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

		lines, updates, err := loadUpdates(file)

		if err != nil {
			return nil, err
		}

		charts = append(charts, &Chart{
			Path:    file,
			Meta:    chart,
			Source:  source,
			Updates: updates,
			Lines:   lines,
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

func loadUpdates(file string) ([]string, []*Update, error) {
	lines, err := readLines(
		path.Join(
			path.Dir(
				file,
			),
			"values.yaml",
		),
	)

	if err != nil {
		return nil, nil, err
	}

	result := make([]*Update, 0)

	for i, line := range lines {
		if matches := updateMatch.FindStringSubmatch(
			line,
		); matches != nil {
			result = append(result, &Update{
				Name:  matches[2],
				Image: matches[3],
				Tag:   matches[1],
				Line:  i,
			})
		}
	}

	return lines, result, nil
}

func readLines(file string) ([]string, error) {
	handle, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer func() { _ = handle.Close() }()

	lines := []string{}
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		lines = append(
			lines,
			scanner.Text(),
		)
	}

	return lines, scanner.Err()
}

func writeLines(file string, lines []string) error {
	handle, err := os.Create(file)

	if err != nil {
		return err
	}

	defer func() { _ = handle.Close() }()

	w := bufio.NewWriter(handle)

	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}

	return w.Flush()
}
