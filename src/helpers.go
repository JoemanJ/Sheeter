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
	"strconv"
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

    case "item_form":
      fmt.Println(form)

      item, err := PTA1.GetItem(form.Get("i_name"))
      if err != nil{
        fmt.Println(err)
      }

      if item.Quantity == 0{
        item, err = PTA1.RegisterItem(form.Get("i_name"), form.Get("i_description"))
        if err != nil{
          fmt.Println(err)
        }
      }

      qtt, err:= strconv.Atoi(form.Get("i_qtt"))
      if err != nil{
        fmt.Println(err)
      }

      factor, err:= strconv.Atoi(form.Get("factor"))
      if err != nil{
        fmt.Println(err)
      }

      for index, i_item := range sheet.Inventory{
        if i_item.Name == item.Name{
          sheet.Inventory[index].Quantity += factor * qtt
          if sheet.Inventory[index].Quantity <= 0{
            sheet.Inventory[index] = sheet.Inventory[len(sheet.Inventory) - 1]
            sheet.Inventory = sheet.Inventory[:len(sheet.Inventory) - 1]
          }
          general.SetJsonData(path, sheet)
          return nil
        }
      }

      sheet.Inventory = append(sheet.Inventory, item)
      sheet.Inventory[len(sheet.Inventory)-1].Quantity = qtt
      general.SetJsonData(path, sheet)

    case "switch_poke":
      pk1, err := strconv.Atoi(form.Get("poke1"))
      if err != nil{
        fmt.Println(err)
      }

      pk2, err := strconv.Atoi(form.Get("poke2"))
      if err != nil{
        fmt.Println(err)
      }

      sheet.PokemonList[pk1], sheet.PokemonList[pk2] = sheet.PokemonList[pk2], sheet.PokemonList[pk1]
      general.SetJsonData(path, sheet)

    case "update":
      atkStage, err := strconv.Atoi(form.Get("atkStage"))
      if err != nil || atkStage < -6 || atkStage > 6{
        fmt.Println(err)
        atkStage = 0
      }

      defStage, err := strconv.Atoi(form.Get("defStage"))
      if err != nil || defStage < -6 || defStage > 6{
        fmt.Println(err)
        defStage = 0
      }

      spatkStage, err := strconv.Atoi(form.Get("spatkStage"))
      if err != nil || spatkStage < -6 || spatkStage > 6{
        fmt.Println(err)
        spatkStage = 0
      }

      spdefStage, err := strconv.Atoi(form.Get("spdefStage"))
      if err != nil || spdefStage < -6 || spdefStage > 6{
        fmt.Println(err)
        spdefStage = 0
      }

      spdStage, err := strconv.Atoi(form.Get("spdStage"))
      if err != nil || spdStage < -6 || spdStage > 6{
        fmt.Println(err)
        spdStage = 0
      }

      class1, err := PTA1.GetTrainerClass(form.Get("class1"))
      if err != nil{
        fmt.Println(err)
      }
      if class1.Name == ""{
        class1 = nil
      }

      class2, err := PTA1.GetTrainerClass(form.Get("class2"))
      if err != nil{
        fmt.Println(err)
      }
      if class2.Name == ""{
        class2 = nil
      }

      class3, err := PTA1.GetTrainerClass(form.Get("class3"))
      if err != nil{
        fmt.Println(err)
      }
      if class3.Name == ""{
        class3 = nil
      }

      class4, err := PTA1.GetTrainerClass(form.Get("class4"))
      if err != nil{
        fmt.Println(err)
      }
      if class4.Name == ""{
        class4 = nil
      }

      sheet.Status.Stages["ATK"] = atkStage
      sheet.Status.Stages["DEF"] = defStage
      sheet.Status.Stages["SPATK"] = spatkStage
      sheet.Status.Stages["SPDEF"] = spdefStage
      sheet.Status.Stages["SPD"] = spdStage
      sheet.Notes = form.Get("notes")

      sheet.CalcStats()

      general.SetJsonData(path, sheet)

    case "allocate_stats":
      hp, err := strconv.Atoi(form.Get("HP"))
      if err != nil{
        fmt.Println(err)
      }

      atk, err := strconv.Atoi(form.Get("ATK"))
      if err != nil{
        fmt.Println(err)
      }

      def, err := strconv.Atoi(form.Get("DEF"))
      if err != nil{
        fmt.Println(err)
      }

      spatk, err := strconv.Atoi(form.Get("SPATK"))
      if err != nil{
        fmt.Println(err)
      }

      spdef, err := strconv.Atoi(form.Get("SPDEF"))
      if err != nil{
        fmt.Println(err)
      }

      spd, err := strconv.Atoi(form.Get("SPD"))
      if err != nil{
        fmt.Println(err)
      }

      vector := map[string]int{"HP": hp, "ATK":atk, "DEF":def, "SPATK":spatk, "SPDEF":spdef, "SPD":spd}

      sheet.AllocateStats(vector)

    default:
      // err = sheet.Update(form)
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
