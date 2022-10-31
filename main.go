/*
Code By:
Diogo "Joe" Delazare Brandão - 2022

This application was made possible by:
The Go proggraming language (Golang) by Google,
The PokeAPI API by ...
*/

package main

import (
	actions "Joe/sheet-hole/pkg/general"
	"Joe/sheet-hole/pkg/pokemon/PTA1"
	"fmt"
)

func main() {
	teste := []string{"string1", "string2", "string3"}

	testMove, err := PTA1.RegisterMove("Nome", "TIPO", "aptitude", teste, 2, actions.CreateDiceSet(2, 6, 6), "Corpo-a-corpo", "Diaria", "Foda-se contest", "Mata o oponente. Mas mata bem morto mesmo.")
	if err != nil {
		panic(err)
	}

	fmt.Println(testMove)
}
