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
	"math/rand"
	"time"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache string //map[string]*template.Template
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// app := &application{
	// 	infoLog:  infoLogger,
	// 	errorLog: errorLogger,
	// }

	//////////////////////////////////////////////////////////////////////////////////////////////
	// ab, _ := PTA1.GetAbility("metabolização")

	sheet, err := PTA1.CreateTrainerSheet("Joe Teste", "Joe", "M", 0, 20, 180, 800, map[string]int{"HP": 10, "ATK": 10, "DEF": 10, "SPATK": 16, "SPDEF": 10, "SPD": 10})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n%+v\n\n", sheet)
}
