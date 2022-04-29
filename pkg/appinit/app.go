package appinit

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"path/filepath"
)

// ChartMetadata is data found in a helm Chart.yaml file
type ChartMetadata struct {
	Name string `json:"Name"`
}

type Application struct {
	root     string
	chartDir string
	name     string
	// filesystem to use. This cannot be changed once
	// an application is created.
	fs afero.Fs
}

func WithFs(f afero.Fs) ApplicationOption {
	return func(a *Application) error {
		a.fs = f
		return nil
	}
}

type ApplicationOption func(*Application) error

// NewApplication validates a Pando application root, and returns a new Application
func NewApplication(dir string, options ...ApplicationOption) (*Application, error) {
	root, err := filepath.Abs(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to determine absolute path for %s", dir)
	}

	//chartDir := filepath.Join(root, "chart")

	a := &Application{
		root: root,
		//chartDir: chartDir,
	}

	//for _, f := range options {
	//	if err = f(a); err != nil {
	//		return nil, errors.Wrapf(err, "options function failed")
	//	}
	//}
	//
	//if a.fs == nil {
	//	a.fs = afero.NewOsFs()
	//}
	//
	//_, err = a.fs.Stat(chartDir)
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		return nil, errors.Errorf("%s is not a Pando application", root)
	//	}
	//
	//	return nil, errors.Wrapf(err, "unable to stat %s", chartDir)
	//}

	//filename := filepath.Join(chartDir, "Chart.yaml")
	//
	//data, err := afero.ReadFile(a.fs, filename)
	//if err != nil {
	//	return nil, errors.Wrapf(err, "failed to read %s", filename)
	//}

	//fmt.Println(data)

	return a, nil
}
