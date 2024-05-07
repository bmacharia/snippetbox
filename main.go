package main

import (
	"log"
	"net/http"
)

// this handler displays the home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Snippetbox"))
}

// this handler displays a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// this handler displays a form for creating a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))

}
func main() {
	mux := http.NewServeMux()

	// this registers the home function as a handler
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
