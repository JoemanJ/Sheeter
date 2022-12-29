package main

import (
	general "Joe/sheeter/pkg/general"
	"Joe/sheeter/pkg/pokemon/PTA1"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (a *application) newAbility(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		f := r.Form
		fmt.Println(f.Get("a_activation"), f.Get("a_description"), f.Get("a_name"))
	}
}

func (a *application) getData(w http.ResponseWriter, r *http.Request) {
	var str string
	_, err := fmt.Sscanf(r.URL.Path, "/data/%s", &str)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if strings.Contains(str, "sheets/") {
		return
	}

	file, err := os.ReadFile("data/" + str + ".json")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(file))
}

func (a *application) newPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			a.serverError(w, err)
			return
		}

		f := r.Form

		switch f.Get("form_name") {
		case "pokemon_form":
			fmt.Println(r.Form)

		case "species_form":
			fmt.Println(r.Form)

		case "ability_form":
			a, err := PTA1.RegisterAbility(f.Get("a_name"), f.Get("a_activation"), f.Get("a_description"))
			fmt.Println(a)
			if err != nil {
				fmt.Println(err)
			}

		case "capacity_form":
			fmt.Println(r.Form)

		}
	}

	a.templateCache["newPokemon.page.html"].Execute(w, nil)
}

func (a *application) generalNew(w http.ResponseWriter, r *http.Request) {
	a.templateCache["new.page.html"].Execute(w, nil)
}

func (a *application) sheet(w http.ResponseWriter, r *http.Request) {
	path, Type, err := general.GetSheetType(1)
	if err != nil {
		fmt.Print(err.Error())
	}
	a.renderSheet(w, r, path, Type)

	return
}

func (a *application) login(w http.ResponseWriter, r *http.Request) {
	a.templateCache["home.page.html"].Execute(w, nil)
	fmt.Println(a.templateCache["home.page.html"])
}
