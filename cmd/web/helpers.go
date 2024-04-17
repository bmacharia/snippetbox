package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// the serverError helper writes log entry at Error level(including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

// the clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in our handlers to send responses like 400 Bad Request
// or 404 Not Found.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// the notHelper function is a convenience wrapprr around clientError which sends a 404 Not Found
// response to the user.// the notHelper function is a convenience wrapprr around clientError which sends a 404 Not Found

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {

	// Retrieve the appropriate template set from the cache based on the page name
	// if no entry in the map then create a new error
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return

	}

	// Write out the provided HTTP status code (200 OK, 400 Bad Request)
	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}

}
