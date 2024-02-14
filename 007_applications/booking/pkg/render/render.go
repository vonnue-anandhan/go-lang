package render

import (
	"booking/pkg/config"
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Renders template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Get template cache from the app config

	// Create template cache
	templateCache, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	// Get requested template from cache
	t, isCached := templateCache[tmpl]

	if !isCached {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)

	err = t.Execute(buff, nil)

	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buff.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all files ending with *.page.tmpl from './templates'
	filePaths, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, filePath := range filePaths {
		fileName := filepath.Base(filePath)
		templateSet, err := template.New(fileName).ParseFiles(filePath)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[fileName] = templateSet
	}

	return myCache, nil
}
