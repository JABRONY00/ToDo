package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type application struct {
	errLg *log.Logger
	infLg *log.Logger
}

func main() {
	logfile, err := os.OpenFile("infoLog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	infLg := log.New(io.MultiWriter(os.Stdout, logfile), "INFO\t", log.Ldate|log.Ltime)
	errLg := log.New(io.MultiWriter(os.Stderr, logfile), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		infLg: infLg,
		errLg: errLg,
	}
	srv := &http.Server{
		Addr:     ":4000",
		ErrorLog: errLg,
		Handler:  app.routes(),
	}
	infLg.Printf("Server Up")
	err = srv.ListenAndServe()
	errLg.Fatal(err)
}
