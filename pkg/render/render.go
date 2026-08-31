package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/richdon/bookings/pkg/config"
	"github.com/richdon/bookings/pkg/models"
)

// app stores the app config to be used in this package
var app *config.AppConfig

// NewRepo setter for the app config
func NewRepo(a *config.AppConfig) {
	app = a
}

func DefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate pulls a template from the cache or creates one on the fly depending on app config
func RenderTemplate(w http.ResponseWriter, name string, td *models.TemplateData) {
	var tc = map[string]*template.Template{}
	if app.UseCache {
		fmt.Println("Using cache")
		tc = app.TemplateCache
	} else {
		fmt.Println("Not Using cache")
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[name]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = DefaultTemplateData(td)

	if err := t.Execute(buf, td); err != nil {
		log.Fatal(err)
	}

	if _, err := buf.WriteTo(w); err != nil {
		log.Fatal(err)
	}

}

// CreateTemplateCache identifies all templates from templates folder and creates a mapping name to template
func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}

	layouts, err := filepath.Glob("/Users/rd/Desktop/bookings/templates/*.layout.html")
	if err != nil {
		return tc, err
	}

	if len(layouts) < 1 {
		return tc, errors.New("could not find required layouts")
	}

	pages, err := filepath.Glob("/Users/rd/Desktop/bookings/templates/*.page.html")
	if err != nil {
		return tc, err
	}

	if len(pages) < 1 {
		return tc, errors.New("could not find required pages")
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return tc, err
		}

		if _, err = ts.ParseGlob("/Users/rd/Desktop/bookings/templates/*.layout.html"); err != nil {
			return tc, err
		}

		tc[name] = ts
	}

	return tc, err
}
