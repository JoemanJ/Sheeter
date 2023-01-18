package main

import (
	general "Joe/sheeter/pkg/general"
	sheeters "Joe/sheeter/pkg/general"
	"Joe/sheeter/pkg/pokemon/PTA1"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "github.com/golang-jwt/jwt"
)

func (a *application) registerNewUser(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST"{
    r.ParseForm()

    user := sheeters.User{
      Username: r.Form.Get("username"),
      Password: r.Form.Get("password"),
      OwnedSheets: map[int]string{},
    }

    userList := map[string]sheeters.User{}

    err := general.GetJsonData(sheeters.USERSDATA, &userList)
    if err != nil{
      fmt.Println(err)
    }

    userList[user.Username] = user

    err = general.SetJsonData(sheeters.USERSDATA, userList)
    if err != nil{
      fmt.Println(err)
    }

    a.login(w, r)
    return
  }

  a.templateCache["register.page.html"].Execute(w, nil)
}

func (a *application) getData(w http.ResponseWriter, r *http.Request) {
	var str string
	_, err := fmt.Sscanf(r.URL.Path, "/data/%s", &str)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if strings.Contains(str, "sheets/") || strings.Contains(str, "users") {
		return
	}

  if strings.HasPrefix(str, "specific"){
    q := r.URL.Query()

    switch q.Get("module"){
    case "PTA2":
      d, err := PTA1.Call(q.Get("type"), q.Get("id"))
      if err != nil{
        fmt.Println("ERRO: ", err)
      }

      data, _ := json.Marshal(d)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprint(w, string(data))
      return
    }

  } else{
    data, err := os.ReadFile("data/" + str + ".json")
    if err != nil {
      w.Write([]byte(err.Error()))
      return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(data))
  }
}

func (a *application) newTrainer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			a.serverError(w, err)
			return
		}

		f := r.Form

    c, err := r.Cookie("sheeter_token")
    if err != nil{
      a.serverError(w, err)
    }

    username := getUsernameFromJWT(c)

    user, err := sheeters.GetUser(username)
    if err != nil{
      a.serverError(w, err)
    }

    dataString := fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s", f["hp"][0], f["atk"][0], f["def"][0], f["spatk"][0], f["spdef"][0], f["spd"][0], f["lvl"][0], f["age"][0], f["height"][0], f["weight"][0])

    var hp, atk, def, spatk, spdef, spd, lvl, age, i_height, i_weight int
    var height, weight float32

    fmt.Sscanf(dataString, "%d;%d;%d;%d;%d;%d;%d;%d;%f;%f", &hp, &atk, &def, &spatk, &spdef, &spd, &lvl, &age, &height, &weight)

    i_height = int(height * 100)
    i_weight = int(weight * 10)

    statMap := map[string]int{"HP":hp, "ATK":atk, "DEF":def, "SPATK":spatk, "SPDEF":spdef, "SPD":spd}


    p, err := PTA1.CreateTrainerSheet(f.Get("name"), user.Username, f.Get("gender"), lvl, age, i_height, i_weight, statMap)
    if err != nil{
      fmt.Println(err)
    }

    fmt.Printf("%+v", user)
    
    m := make(map[int]string)
    m = user.OwnedSheets
    m[p.Id] = p.Name

    user.OwnedSheets = m
    sheeters.SetUser(user)

    http.Redirect(w, r, "/sheet/?id="+strconv.Itoa(p.Id), http.StatusSeeOther)
	}

	err := a.templateCache["newTrainer.page.html"].Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *application) newPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			a.serverError(w, err)
			return
		}

		f := r.Form
    
    c, err := r.Cookie("sheeter_token")
    if err != nil{
      a.serverError(w, err)
    }

    username := getUsernameFromJWT(c)

    user, err := sheeters.GetUser(username)
    if err != nil{
      a.serverError(w, err)
    }

		switch f.Get("form_name") {
		case "pokemon_form":
			a, err := PTA1.GetAbility(f.Get("ability"))
			if err != nil {
				fmt.Println(err)
			}
			abilities := []*PTA1.PokemonAbility{a}

			lvl, err := strconv.Atoi(f.Get("lvl"))
			if err != nil {
				fmt.Println(err)
			}

			poke, err := PTA1.CreatePokemonSheet(f.Get("nickname"), f.Get("species"), f.Get("gender"), f.Get("nature"), abilities, lvl)
			if err != nil {
				fmt.Println(err)
			}
      
      if f.Get("referrer_sheet") != ""{
        refs := &PTA1.TrainerSheet{}

        err = general.GetJsonData("data/sheets/" + f.Get("referrer_sheet") + "_0.json", refs)
        if err != nil{
          fmt.Println(err)
          return
        }

        if !refs.CaughtPokemon[poke.Species.Number-1]{
          refs.TotalCaughtPokemon++
        }
        refs.CaughtPokemon[poke.Species.Number-1] = true

        if !refs.SeenPokemon[poke.Species.Number-1]{
          refs.TotalSeenPokemon++
        }
        refs.SeenPokemon[poke.Species.Number-1] = true

        refs.PokemonList = append(refs.PokemonList, poke)
        general.SetJsonData("data/sheets/" + f.Get("referrer_sheet") + "_0.json", refs)

        user.OwnedSheets[poke.Id] = poke.Nick
        general.SetUser(user)

        http.Redirect(w, r, "/sheet/?id="+strconv.Itoa(poke.Id), http.StatusSeeOther)
      }

		case "species_form":
			var abilities []*PTA1.PokemonAbility
			var high_abilities []*PTA1.PokemonAbility
			var capacities []*PTA1.Capacity
			basicCapacities := [3]int{0, 0, 0}
			movement := map[string]int{"land": 0, "surface": 0, "underwater": 0, "burrow": 0, "fly": 0}

			for k, v := range f {
				if strings.Contains(k, "c_") {
					c, err := PTA1.GetCapacity(strings.Replace(k, "c_", "", 1))
					if err != nil {
						fmt.Println(err)
					}

					capacities = append(capacities, c)
					continue
				}

				if strings.Contains(k, "ha_") {
					aux := strings.Replace(k, "ha_", "", 1)

					ha, err := PTA1.GetAbility(aux)
					if err != nil {
						fmt.Println(err)
					}

					if err != nil {
						fmt.Println(err)
					}

					high_abilities = append(high_abilities, ha)
					continue
				}

				if strings.Contains(k, "m_") && k != "form_name" { //I didn't thing the names through beforehand...
					n, err := strconv.Atoi(v[0])
					if err != nil {
						fmt.Println(err)
					}

					aux := strings.Replace(k, "m_", "", 1)

					movement[aux] = n
					continue
				}

				if k == "s_inteligence" || k == "s_jump" || k == "s_strength" {
					n, err := strconv.Atoi(v[0])
					if err != nil {
						fmt.Println(err)
					}

					switch k {
					case "s_inteligence":
						basicCapacities[2] = n
					case "s_jump":
						basicCapacities[1] = n
					case "s_strength":
						basicCapacities[0] = n
					}

					continue
				}

				if strings.Contains(k, "a_") {
					aux := strings.Replace(k, "a_", "", 1)

					a, err := PTA1.GetAbility(aux)
					if err != nil {
						fmt.Println(err)
					}

					abilities = append(abilities, a)
					continue
				}
			}

			_, err := PTA1.RegisterSpecies(string(f.Get("s_species_name")), string(f.Get("s_diet")), basicCapacities, capacities, abilities, high_abilities, movement)
			if err != nil {
				fmt.Println(err)
			}

		case "ability_form":
			a, err := PTA1.RegisterAbility(f.Get("a_name"), f.Get("a_activation"), f.Get("a_description"))
			fmt.Println(a)
			if err != nil {
				fmt.Println(err)
			}

		case "capacity_form":
			_, err := PTA1.RegisterCapacity(f.Get("c_name"), f.Get("c_description"))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	a.templateCache["newPokemon.page.html"].Execute(w, nil)
}

func (a *application) generalNew(w http.ResponseWriter, r *http.Request) {
	a.templateCache["new.page.html"].Execute(w, nil)
}

func (a *application) sheet(w http.ResponseWriter, r *http.Request) {
  id, err:= strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil{
    a.generalNew(w, r)
    return
  }

  c, err := r.Cookie("sheeter_token")
  if err != nil{
    fmt.Println(err)
  }

  username := getUsernameFromJWT(c)

  user, err := sheeters.GetUser(username)
  if err != nil{
    a.serverError(w, err)
  }

  if _, ok := user.OwnedSheets[id]; !ok{
    http.Redirect(w, r, "/login", http.StatusUnauthorized)
    return
  }

	path, Type, err := general.GetSheetType(id)
	if err != nil {
		fmt.Print(err.Error())
	}

  if r.Method == "POST"{
    err = r.ParseForm()
    if err != nil{
      fmt.Println(err)
    }

    if err == nil{
      f := r.Form 
      a.handleSheetUpdates(path, Type, f)
    }
    
  }

	a.renderSheet(w, r, path, Type)

	return
}

func (a *application) login(w http.ResponseWriter, r *http.Request) {
  c, err := r.Cookie("sheeter_token")
  if err == nil && validateJWT_(c){
    a.generalNew(w, r)
    return
  }

  r.ParseForm()
  form := r.Form

  if form.Has("username") && form.Has("password"){
    ul := map[string]sheeters.User{}

    err := general.GetJsonData(sheeters.USERSDATA, &ul)
    if err != nil{
      fmt.Println(err)
    }

    usr, ok := ul[form.Get("username")]

    if !ok || usr.Password != form.Get("password"){
      r.Header.Add("error", "invalid_combination")
      a.templateCache["login.page.html"].Execute(w, nil)
      return
    }

    jwt, err := generateJWT(usr.Username)
    if err != nil{
      r.Header.Add("error", "token_generation_failed")
      fmt.Println(err)
      a.templateCache["login.page.html"].Execute(w, nil)
      return
    }

    http.SetCookie(w, &http.Cookie{
      Name: "sheeter_token",
      Value: jwt,
    })

    a.generalNew(w, r)
    return
  }

  a.templateCache["login.page.html"].Execute(w, nil)
  return

}

func (a *application) account(w http.ResponseWriter, r *http.Request){
  t, err := sheeters.RenderVolatile("account.page.html", "ui/html")
  if err != nil{
    fmt.Println(err)
    return
  }

  c, err := r.Cookie("sheeter_token")
  if err != nil{
    http.Redirect(w, r, "/login", http.StatusUnauthorized)
  }

  username := getUsernameFromJWT(c)

  user, err := sheeters.GetUser(username)
  if err != nil{
    http.Redirect(w, r, "/login", http.StatusUnauthorized)
  }
  fmt.Printf("%+v\n", user)

  t.Execute(w, *user)
}

func (a *application) logout(w http.ResponseWriter, r *http.Request){
    http.SetCookie(w, &http.Cookie{
      Name: "sheeter_token",
      Value: "invalid",
      Path: "/",
    })

    http.Redirect(w, r, "/login", http.StatusOK)
}
