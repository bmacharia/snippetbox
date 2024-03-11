package main

import (
	"log"
	"net/http"
)

// home is a handler function that writes a response to the http.ResponseWriter.
// the home handler writes a byte slice containing "Hello from Snippetbox".
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// This is a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// use the r.Methiod to check whether or not the request to the server is a POST request or not
	if r.Method != "POST" {
		// if the request is not a POST request
		// use th WriteHeader(405) method to send a 405 status code and the corresponding description to the client
		// the code below is not executed if the request is not a POST request
		// use the w.Writeheader() method to send a non-200 status code

		w.Write([]byte("Method Not Allowed\n"))
		return

	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Register the two new handler functions and corresponding URL patterns with the servemux, using the HandleFunc() method. This means that the home handler will be used when a request is made to the / URL pattern, and the snippet handler will be used when a request is made to the /snippet URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server. We pass in two parameters: the TCP network address to listen on (in this case ":4000") and the servemux we just created.
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
