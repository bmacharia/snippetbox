package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"snippetbox.bmacharia/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	//call the newTemplateData helper function to get templateData struct
	// containing the current year, and it to the snippets slice
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// pass the data to the render functon
	app.render(w, r, http.StatusOK, "home.tmpl", data)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	// use of the new render helper function
	app.render(w, r, http.StatusOK, "view.tmpl", data)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm() adds anu data to POST request bodies to the r.PostForm map
	//if there any errors, we use our app.clientError helper to send a 400 Bad Request
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//use r.PostForm() method to retrieve the title and content
	// from the r.PostForm map
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// manually convert form data into an integer using the strcoonv.Atoi(), and send a  400 Bad Request if it fails
	//
	//expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//
	//}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
