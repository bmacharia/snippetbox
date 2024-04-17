// In a real world application multiple pieces of dynamic data
// tht I want to display in the same page
// a lightweight way to pass data to the template is to use a dynamic struct which
// acts like a single structure for the data

package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.bmacharia/internal/models"
)

// define a templateData type to act as a holding structure for any dynamic data that I want to pass to my HTML templates
type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	//Intialized a new map to act as the cache
	cache := map[string]*template.Template{}

	// Use the filepath.Glob function to get a slice of all the filepaths
	// that match the pattern "./ui/html/*.tmpl"
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// use a range loop t loop through the filepaths one by one
	for _, page := range pages {
		//extract the base name of the file and assign it to the name variable
		name := filepath.Base(page)

		// Pares the base template file into a tepmlate set
		ts, err := template.ParseFiles("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParsGlob() * on this template set to add any partials
		ts, err = ts.ParseGlob("./ui/html/pages/*.tmpl")
		if err != nil {
			return nil, err

		}

		// call ParseFiles() *on this template ser to add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
