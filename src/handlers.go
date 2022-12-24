package main

import (
	general "Joe/sheeter/pkg/general"
	"fmt"
	"net/http"
)

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
