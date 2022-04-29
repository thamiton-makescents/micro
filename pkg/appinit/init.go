package appinit

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed python
var assets embed.FS

var Boxes = map[string]string{
	"python": "python",
}

type AppInitializer struct {
	dir string
	fs  afero.Fs
	out io.Writer
}

func NewAppInitializer(dir string, fs afero.Fs, out io.Writer) *AppInitializer {
	return &AppInitializer{
		dir: dir,
		fs:  fs,
		out: out,
	}
}

func (i *AppInitializer) Run(name string, lang string) (*Application, error) {
	if err := i.fs.MkdirAll(i.dir, 0755); err != nil {
		return nil, fmt.Errorf("unable to create application directory")
	}

	box, ok := Boxes[strings.ToLower(lang)]
	if !ok {
		return nil, fmt.Errorf("failed to get language box for %s", lang)
	}

	fileList := make([]string, 0)
	if err := fs.WalkDir(assets, box, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		fileList = append(fileList, strings.TrimPrefix(path, d.Name()+"/"))
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to traverse directory for %s", box)
	}

	if err := i.writeStaticAssets(name, fileList); err != nil {
		return nil, fmt.Errorf("failed to build the static assets for %s: %w", box, err)
	}

	return NewApplication(i.dir, WithFs(i.fs))
}

type templateParams struct {
	Name      string
	TitleName string
}

// writeStaticAssets() parses through the static assets and
// writes them as files in the new application's directory
func (i *AppInitializer) writeStaticAssets(name string, fileList []string) error {
	funcMap := template.FuncMap{
		"Capitalize": strings.Title,
	}

	for _, n := range fileList {
		raw, err := assets.ReadFile(n)
		if err != nil {
			return errors.Wrapf(err, "could not get data for asset %s as a string", n)
		}

		tplVars := templateParams{
			Name:      name,
			TitleName: strings.Title(name),
		}

		var d bytes.Buffer
		rawString := string(raw)
		tmpl, err := template.New(n).Funcs(funcMap).Parse(rawString)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&d, tplVars)
		if err != nil {
			return err
		}

		trimmed := strings.Join(strings.Split(n, "/")[1:], "/")

		// if file is a bash script, make it executable
		if strings.HasPrefix(rawString, "#!") {
			if err := i.writeFile(trimmed, d.Bytes(), 0755); err != nil {
				return errors.Wrapf(err, "could not save asset %s", n)
			}
		} else {
			if err := i.writeFile(trimmed, d.Bytes(), 0644); err != nil {
				return errors.Wrapf(err, "could not save asset %s", n)
			}
		}
	}
	return nil
}

// writeFile() writes a file to a path under the new application's root, creating
// any intermediate directories as needed.
func (i *AppInitializer) writeFile(path string, data []byte, mode os.FileMode) error {
	fullpath := filepath.Join(i.dir, path)

	exists, err := afero.Exists(i.fs, fullpath)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("cannot overwrite existing file")
	}

	if err := i.fs.MkdirAll(filepath.Dir(fullpath), 0755); err != nil {
		return err
	}

	// We'll log writing .notgo files once we rename them later
	if !strings.HasSuffix(fullpath, ".notgo") {
		fmt.Fprintf(i.out, "Creating file %s\n", fullpath)
	}

	return afero.WriteFile(i.fs, fullpath, data, mode)
}
