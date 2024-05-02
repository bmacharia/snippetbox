package main

import (
	"html/template"
	"net/url"
	"path/filepath"
	"time"

	"snippetbox.bmacharia/internal/models"
)

// define a templateData type to act as a holding structure for any dynamic data that I want to pass to my HTML templates
type templateData struct {
	CurrentYear int
	FormData    url.Values
	FormErrors  map[string]string
	Snippet     models.Snippet
	Snippets    []models.Snippet
	Form        any
}

// create a function that returns the date in a human friendly format
// this will represent a time.Time object
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// intialize a template.FuncMap object and then populate it with the humanDate function
// this is a string-keyed map which acts as a loolkup between the names of the functions that I want to use in my templates
var functions = template.FuncMap{
	"humanDate": humanDate,
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

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParsGlob() * on this template set to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
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
