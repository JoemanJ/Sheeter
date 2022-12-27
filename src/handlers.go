package main

import (
	general "Joe/sheeter/pkg/general"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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
