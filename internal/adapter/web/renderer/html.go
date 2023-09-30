package renderer

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/mapper"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type renderer struct {
	templates map[string]*template.Template
}

func New() port.Renderer {
	return &renderer{templates: make(map[string]*template.Template)}
}

func (r renderer) Load(path string) error {
	var (
		pages   []string
		layouts []string
		err     error
	)

	pages, err = r.LoadPageFiles(path)
	if err != nil {
		return err
	}

	layouts, err = r.LoadLayoutFiles(path)
	if err != nil {
		return err
	}

	for _, page := range pages {
		var (
			pageName string
			tmpl     *template.Template
		)

		pageName = filepath.Base(page)

		tmpl = template.New(pageName)
		tmpl, err = tmpl.ParseFiles(layouts...)
		if err != nil {
			return err
		}
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return err
		}

		r.templates[pageName] = tmpl
		fmt.Println(r.templates[pageName], "\t", tmpl)
	}
	return nil
}

func (r renderer) Render(page string, wr io.Writer, data ...interface{}) error {
	var (
		tmpl *template.Template
		err  error
	)

	tmpl = r.templates[page]
	if tmpl == nil {
		return fmt.Errorf("cannot find page with name %s", page)
	}

	err = tmpl.Execute(wr, data)
	if err != nil {
		return err
	}

	return nil
}

func (r renderer) LoadPageFiles(path string) ([]string, error) {
	var (
		paths []string
		err   error
	)

	paths, err = r.LoadFilesContaining(path, "page")
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (r renderer) LoadLayoutFiles(path string) ([]string, error) {
	var (
		paths []string
		err   error
	)

	paths, err = r.LoadFilesContaining(path, "layout")
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (r renderer) LoadFilesContaining(path string, contains string) ([]string, error) {
	var (
		filePaths map[string]string
		paths     = make([]string, 0)
		err       error
	)

	filePaths, err = mapper.FilesToPaths(path)
	if err != nil {
		return nil, err
	}

	for name, path := range filePaths {
		if !strings.Contains(name, contains) {
			continue
		}
		paths = append(paths, path)
	}

	return paths, nil
}
