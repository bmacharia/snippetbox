package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
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

	// a new byte buffer
	buf := new(bytes.Buffer)

	// write the template into the buffer instead of the http.ResponseWriter
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// if the template is written to the buffer without error
	// write the http status code to the http.ResponseWriter
	w.WriteHeader(status)

	// write the contens of the buffer to the http.ResponseWriter
	buf.WriteTo(w)

}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil

}
