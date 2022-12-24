package main

import (
	general "Joe/sheeter/pkg/general"
	"Joe/sheeter/pkg/pokemon/PTA1"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime/debug"
)

func (a *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	a.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (a *application) newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/*.page.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (a *application) renderSheet(w http.ResponseWriter, r *http.Request, path string, Type int) {
	switch Type {
	case 0:
		var sheet *PTA1.TrainerSheet

		general.GetJsonData(path, sheet)

		err := sheet.Render(w)
		if err != nil {
			a.serverError(w, err)
		}

	case 1:
		var sheet *PTA1.PokemonSheet

		general.GetJsonData(path, sheet)
		err := sheet.Render(w)
		if err != nil {
			a.serverError(w, err)
		}

	default:
		w.Write([]byte("Unknown sheet type"))
	}

}
