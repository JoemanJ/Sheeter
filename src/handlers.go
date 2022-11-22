package main

import (
	general "Joe/sheet-hole/pkg/general"
	"Joe/sheet-hole/pkg/pokemon/PTA1"
	"fmt"
	"net/http"
)

func (a *application) sheet(w http.ResponseWriter, r *http.Request) {
	sheet, err := general.GetSheet(0)
	if err != nil {
		fmt.Print(err.Error())
	}

	ts, err := sheet.SheetBody()
	if err != nil {
		a.serverError(w, err)
	}

	ts.Execute(w, "")
}

func (a *application) login(w http.ResponseWriter, r *http.Request) {
	_, err := PTA1.CreateTrainerSheet("joe", "Joe", "M", 0, 12, 1, 1, map[string]int{"HP": 10, "ATK": 10, "DEF": 10, "SPATK": 10, "SPDEF": 10, "SPD": 10})
	if err != nil {
		a.serverError(w, err)
	}
}
