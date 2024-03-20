package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// slog.New() function intializes a new structured logger
	// which writes to the standard output stream and uses default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// the Info() method to log the starting server message Info severity
	// alonmg with the listen and server attribute
	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	// the Error() method to log any error message returned by listen and serve
	// http.ListenAndServe() at Error severity
	// and then call the os.Exit(1) to terminate the application with exit code 1
	logger.Error(err.Error())

	os.Exit(1)
}
