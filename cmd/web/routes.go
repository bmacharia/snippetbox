package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// create a middlewate chain containing our standard middleware
	//which will be used for every request our application receives
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// return the 'standard middleware chain' followed by the servemux
	return standard.Then(mux)

}
