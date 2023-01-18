/*
Code By:
Diogo "Joe" Delazare Brandão - 2022

This application was made possible by:
The Go proggraming language (Golang) by Google,
The PokeAPI API by ...
*/

package main

import (
	// sheeters "Joe/sheeter/pkg/general"
	// "Joe/sheeter/pkg/pokemon/PTA1"
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

const JWTKEY = "CHANGETHIS!"

const SHEETSPATH = "data/sheet/"

func main() {
	_, err := os.Stat("./data")
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir("data", 0755)
	}

	_, err = os.Stat("./data/sheets")
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir("data/sheets", 0755)
	}

  // sheeters.SetJsonData("./data/PTA1/trainerLvlTable.json", PTA1.TRAINERLVLTABLE)

	rand.Seed(time.Now().UnixNano())

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog, templateCache: map[string]*template.Template{},
	}

	cache, err := app.newTemplateCache("./ui/html")

	if err != nil {
		panic(err)
	}

	app.templateCache = cache

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.login)
	mux.HandleFunc("/new/", app.generalNew)
	mux.HandleFunc("/new/pokemon", app.newPokemon)
	mux.HandleFunc("/new/trainer", app.newTrainer)
	mux.HandleFunc("/data/", app.getData)
	mux.Handle("/sheet/", validateJWT(http.HandlerFunc(app.sheet)))
  mux.Handle("/account/", validateJWT(http.HandlerFunc(app.account)))
	mux.HandleFunc("/register", app.registerNewUser)
	mux.HandleFunc("/login", app.login)
  mux.HandleFunc("/logout/", app.logout)
	// sheet, err := PTA1.CreateSampleTrainerSheet()
	if err != nil {
		panic(err)
	}

	// fmt.Println(sheet)
	fmt.Printf("starting server on port %s\n", PORT)
	err = http.ListenAndServe(PORT, mux)
	if err != nil {
		println(err.Error())
	}
}
