/*
Code By:
Diogo "Joe" Delazare Brandão - 2022

This application was made possible by:
The Go proggraming language (Golang) by Google,
The PokeAPI API by ...
*/

package main

import (
	sheeters "Joe/sheet-hole/pkg/general"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

const PORT string = ":4000"

func main() {
	_, err := os.Stat("./data")
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir("data", 0755)
	}

	_, err = os.Stat("./data/sheets")
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir("data", 0755)
	}

	rand.Seed(time.Now().UnixNano())

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: map[string]*template.Template{},
	}

	cache, err := app.newTemplateCache("./ui/html")
	if err != nil {
		panic(err)
	}

	app.templateCache = cache

	mux := http.NewServeMux()

	err = sheeters.SetRD("sheetCount", "0")
	if err != nil {
		app.errorLog.Panic(err)
	}

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/new/", app.login)
	mux.HandleFunc("/sheet", app.sheet)

	fmt.Printf("starting server on port %s\n", PORT)
	err = http.ListenAndServe(PORT, mux)
	if err != nil {
		println(err.Error())
	}
}
