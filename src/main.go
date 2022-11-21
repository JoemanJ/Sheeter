/*
Code By:
Diogo "Joe" Delazare Brandão - 2022

This application was made possible by:
The Go proggraming language (Golang) by Google,
The PokeAPI API by ...
*/

package main

import (
	"log"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache string //map[string]*template.Template
}

const PORT string = ":4000"

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/", Sheet)

// 	fmt.Printf("starting server on port %s\n", PORT)
// 	err := http.ListenAndServe(PORT, mux)
// 	if err != nil {
// 		println(err.Error())
// 	}
// }
