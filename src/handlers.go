package main

import (
	general "Joe/sheeter/pkg/general"
	"Joe/sheeter/pkg/pokemon/PTA1"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

func (a *application) newTrainer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			a.serverError(w, err)
			return
		}

		f := r.Form

		switch f.Get("form_name") {
		case "":
			// PTA1.crea
		}
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

		switch f.Get("form_name") {
		case "pokemon_form":
			fmt.Println(r.Form)
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
			fmt.Println(poke)

		case "species_form":
			var abilities []*PTA1.PokemonAbility
			var high_abilities []*PTA1.PokemonAbility
			var capacities []*PTA1.Capacity
			basicCapacities := [3]int{0, 0, 0}
			movement := map[string]int{"land": 0, "surface": 0, "underwater": 0, "burrow": 0, "fly": 0}
			fmt.Println(f)

			for k, v := range f {
				if strings.Contains(k, "c_") {
					c, err := PTA1.GetCapacity(k)
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

					fmt.Println(ha, k, aux)
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
			a, err := PTA1.RegisterCapacity(f.Get("c_name"), f.Get("c_description"))
			fmt.Println(a)
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
