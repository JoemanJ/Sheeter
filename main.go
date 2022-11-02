/*
Code By:
Diogo "Joe" Delazare Brandão - 2022

This application was made possible by:
The Go proggraming language (Golang) by Google,
The PokeAPI API by ...
*/

package main

import (
	"Joe/sheet-hole/pkg/pokemon/PTA1"
	"fmt"
	"log"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache string //map[string]*template.Template
}

func main() {
	// infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// app := &application{
	// 	infoLog:  infoLogger,
	// 	errorLog: errorLogger,
	// }

	// PTA1.RegisterAbility("Metabolização", "Constante", "Imunidade a Venenos")
	// PTA1.RegisterAbility("Alucinógeno", "Constante", "quando este pokémon recebe Veneno, eleve duas Fases do Ataque dele. Se ele for curado da Condição, ele perde as duas Fases do Ataque.")
	PTA1.RegisterItem("test item", "This item tests if item registration works")

	i, _ := PTA1.GetItem("test item")

	fmt.Printf("%+v\n", i)
}
