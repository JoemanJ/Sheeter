package main

import (
	general "Joe/sheet-hole/pkg/general"
	"Joe/sheet-hole/pkg/pokemon/PTA1"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *application) sheet(w http.ResponseWriter, r *http.Request) {
	ts := a.templateCache["sheet.page.html"]

	sheet, err := general.GetSheet(0)
	if err != nil {
		fmt.Print(err.Error())
	}

	switch sheet.Type {
	case "PTA1_trainerSheet":
		data, err := json.MarshalIndent(sheet.Data, "", "  ")
		if err != nil {
			a.serverError(w, err)
			return
		}

		var s PTA1.TrainerSheet

		err = json.Unmarshal(data, &s)
		if err != nil {
			a.serverError(w, err)
			return

		}

		err = s.Render(w, ts)
		if err != nil {
			a.serverError(w, err)
			return
		}

		w.Write([]byte("Unknown sheet type"))
		return
	}

	w.Write([]byte("Something went wrong!"))
	return
}

func (a *application) login(w http.ResponseWriter, r *http.Request) {
	a.templateCache["home.page.html"].Execute(w, nil)
	fmt.Println(a.templateCache["home.page.html"])
}
