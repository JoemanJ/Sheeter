package main

import (
	general "Joe/sheeter/pkg/general"
	"Joe/sheeter/pkg/pokemon/PTA1"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"
	"runtime/debug"
	"strings"
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
		sheet := &PTA1.TrainerSheet{}

		general.GetJsonData(path, sheet)
		err := sheet.Render(w)
		if err != nil {
			a.serverError(w, err)
		}

	case 1:
		sheet := &PTA1.PokemonSheet{}

		err := general.GetJsonData(path, sheet)
		if err != nil {
			a.serverError(w, err)
		}

		err = sheet.Render(w)
		if err != nil {
			a.serverError(w, err)
		}

	default:
		w.Write([]byte("Unknown sheet type"))
	}

}

func (a *application) handleSheetUpdates(path string, Type int, form url.Values) error{
  switch Type{
	case 0:
		sheet := &PTA1.TrainerSheet{}

    err := general.GetJsonData(path, sheet)
    if err != nil{
      return err
    }

    switch form.Get("form_name") {
    case "class_form":
      
      fmt.Println(form, "\n\n", "AQUI")
      basicTalents := [2]*PTA1.TrainerTalent{}
      otherTalents := []*PTA1.TrainerTalent{}
      expertises := []*PTA1.Expertise{}

      bt1, err := PTA1.GetTrainerTalent(form.Get("class_basic_talent1"))
      if err != nil{
        fmt.Println(err)
      }

      basicTalents[0] = bt1

      bt2, err := PTA1.GetTrainerTalent(form.Get("class_basic_talent2"))
      if err != nil{
        fmt.Println(err)
      }

      basicTalents[1] = bt2

      for data := range form{
        if strings.HasPrefix(data, "t_"){
          t, err := PTA1.GetTrainerTalent(strings.TrimPrefix(data, "t_"))
          if err != nil{
            fmt.Println(err)
            continue
          }

          otherTalents = append(otherTalents, t)
        }

        if strings.HasPrefix(data, "e_"){
          e, err := PTA1.GetExpertise(strings.TrimPrefix(data, "e_"))
          if err != nil{
            fmt.Println(err)
            continue
          }

          fmt.Println(e)

          expertises = append(expertises, e)
        }
      }

      class, err := PTA1.RegisterTrainerClass(form.Get("class_name"), form.Get("class_description"), form.Get("class_parent"), basicTalents, otherTalents, expertises, form.Get("class_requirements"))
      if err != nil{
        fmt.Println(err)
      }

      fmt.Println(class)
      
    case "talent_form":

      fmt.Println(form)
      t, err:= PTA1.RegisterTrainerTalent(form.Get("talent_name"), false, form.Get("talent_requirements"), form.Get("talent_frequency"), form.Get("talent_target"), form.Get("talent_description"), form.Has("talent_continuous"), form.Has("talent_standart"), form.Has("talent_free"), form.Has("talent_interrupt"), form.Has("talent_extended"), form.Has("talent_legal"))
      if err!= nil{
        return err
      }
      fmt.Println(t)
      return nil

    case "expertise_form":
      fmt.Println(form)

      ex, err := PTA1.RegisterExpertise(form.Get("expertise_name"), form.Get("expertise_stat"), form.Get("expertise_description"))
      if err != nil{
        return err
      }

      fmt.Println(ex)
      return nil

    default:
      err = sheet.Update(form)
      if err != nil{
        return err
      }
    }

    return nil

	case 1:
		sheet := &PTA1.PokemonSheet{}

		err := general.GetJsonData(path, sheet)
		if err != nil {
			return err
		}


	default:
		return errors.New("Unknown sheet type")
	}

  return nil
}
