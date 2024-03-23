package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// slog.New() function intializes a new structured logger
	// which writes to the standard output stream and uses default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Here I will intilize a new instance of the application struct
	// the application struct will containt the dependencies for the application
	// the dependencies that we have thus far
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	// the Error() method to log any error message returned by listen and serve
	// http.ListenAndServe() at Error severity
	// and then call the os.Exit(1) to terminate the application with exit code 1
	logger.Error(err.Error())

	os.Exit(1)
}
