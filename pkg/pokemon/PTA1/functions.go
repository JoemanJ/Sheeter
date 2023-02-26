package PTA1

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"

	general "Joe/sheeter/pkg/general"
)

const ABILITYDATA string = "./data/PTA1/abilityData.json"
const MOVEDATA string = "./data/PTA1/moveData.json"
const SPECIESDATA string = "./data/PTA1/speciesData.json"
const ITEMDATA string = "./data/PTA1/itemData.json"
const TALENTDATA string = "./data/PTA1/talentData.json"
const CLASSDATA string = "./data/PTA1/classData.json"
const EXPERTISEDATA string = "./data/PTA1/expertiseData.json"
const CAPACITYDATA string = "./data/PTA1/capacityData.json"

var GETFUNCMAP map[string]any = map[string]any {
  "ability":GetAbility,
  "move": GetMove,
  "species":GetSpecies,
  "item":GetItem,
  "trainerTalent":GetTrainerTalent,
  "trainerClass":GetTrainerClass,
  "expertise":GetExpertise,
  "capacity":GetCapacity,
}

func Call(funcName, id string)(any, error){
  f := reflect.ValueOf(GETFUNCMAP[funcName])
  var res []reflect.Value
  var in []reflect.Value
  in = append(in, reflect.ValueOf(id))
  res = f.Call(in)
  if !res[1].IsZero(){
    return nil, res[1].Interface().(error)
  }

  return res[0].Interface(), nil  
}

func RegisterAbility(name string, activation string, description string) (*PokemonAbility, error) {
	newAbility := &PokemonAbility{
		Name:        strings.Title(name),
		Activation:  activation,
		Description: description,
	}

	var abilities map[string]PokemonAbility

	err := general.GetJsonData(ABILITYDATA, &abilities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ABILITYDATA, err.Error())
		return nil, errors.New(s)
	}

	abilities[strings.ToLower(name)] = *newAbility

	err = general.SetJsonData(ABILITYDATA, abilities)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", ABILITYDATA, err.Error())
		return nil, errors.New(s)
	}

	return newAbility, nil
}

func GetAbility(name string) (*PokemonAbility, error) {
	var abilities map[string]PokemonAbility
	err := general.GetJsonData(ABILITYDATA, &abilities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ABILITYDATA, err.Error())
		return &PokemonAbility{}, errors.New(s)
	}

	PA := abilities[strings.ToLower(name)]

	return &PA, nil
}

func RegisterMove(name, Type, aptitude string, descriptors []string, accDiff int, dice *general.DiceSet, reach, frequency, contests, effect string) (*PokemonMove, error) {
	//ADD NEW MOVE TO MAP
	newMove := &PokemonMove{
		Name:        strings.Title(name),
		Type:        Type,
		Aptitude:    aptitude,
		Descriptors: descriptors,
		AccDiff:     accDiff,
		Damage:      dice,
		Reach:       reach,
		Frequency:   frequency,
		Contests:    contests,
		Effect:      effect,
	}

	var moves map[string]PokemonMove

	err := general.GetJsonData(MOVEDATA, &moves)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", MOVEDATA, err.Error())
		return nil, errors.New(s)
	}

	moves[strings.ToLower(name)] = *newMove

	err = general.SetJsonData(MOVEDATA, moves)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", MOVEDATA, err.Error())
		return nil, errors.New(s)
	}

	return newMove, nil
}

func GetMove(name string) (*PokemonMove, error) {
	var moves map[string]PokemonMove
	err := general.GetJsonData(MOVEDATA, &moves)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", MOVEDATA, err.Error())
		return &PokemonMove{}, errors.New(s)
	}

	PM := moves[strings.ToLower(name)]

	return &PM, nil
}

func RegisterSpecies(name string, diet string, capacities [3]int, others []*Capacity, abilities []*PokemonAbility, highAbilities []*PokemonAbility, movement map[string]int) (*PokemonSpecies, error) {
	pokemon, err := pokeapi.Pokemon(strings.ToLower(name))
	if err != nil {
		s := fmt.Sprintf("Error getting data from PokeAPI:\n%v", err)
		return nil, errors.New(s)
	}

	typesData := pokemon.Types

	var types []string

	for _, t := range typesData {
		types = append(types, t.Type.Name)
	}

	capacitieTable := newCapacityTable(capacities, others)

	newSpecies := &PokemonSpecies{
		Number: pokemon.ID,
		Name:   strings.Title(pokemon.Species.Name),

		Sprite: pokemon.Sprites.FrontDefault,

		Type:       types,
		Diet:       diet,
		Capacities: capacitieTable,
		Movement:   movement,

		AverageHeight: pokemon.Height * 10,
		AverageWeight: pokemon.Weight,
		BaseStats:     map[string]int{"HP": (pokemon.Stats[0].BaseStat + 5) / 10, "ATK": (pokemon.Stats[1].BaseStat + 5) / 10, "DEF": (pokemon.Stats[2].BaseStat + 5) / 10, "SPATK": (pokemon.Stats[3].BaseStat + 5) / 10, "SPDEF": (pokemon.Stats[4].BaseStat + 5) / 10, "SPD": (pokemon.Stats[5].BaseStat + 5) / 10},

		Abilities:     abilities,
		HighAbilities: highAbilities,
	}

	var species map[string]PokemonSpecies

	err = general.GetJsonData(SPECIESDATA, &species)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", SPECIESDATA, err.Error())
		return nil, errors.New(s)
	}

	species[strings.ToLower(name)] = *newSpecies

	err = general.SetJsonData(SPECIESDATA, species)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", SPECIESDATA, err.Error())
		return nil, errors.New(s)
	}

	return newSpecies, nil
}

func GetSpecies(name string) (*PokemonSpecies, error) {
	var species map[string]PokemonSpecies
	err := general.GetJsonData(SPECIESDATA, &species)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", SPECIESDATA, err.Error())
		return &PokemonSpecies{}, errors.New(s)
	}

	PS := species[strings.ToLower(name)]

	return &PS, nil
}

func RegisterItem(name string, description string) (*Item, error) {
	i := &Item{
		Quantity: 0,

		Name:        strings.Title(name),
		Description: description,
	}

	var items map[string]Item

	err := general.GetJsonData(ITEMDATA, &items)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ITEMDATA, err.Error())
		return nil, errors.New(s)
	}

	items[strings.ToLower(name)] = *i

	err = general.SetJsonData(ITEMDATA, items)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", ITEMDATA, err.Error())
		return nil, errors.New(s)
	}

	return i, nil
}

func GetItem(name string) (*Item, error) {
	var items map[string]Item
	err := general.GetJsonData(ITEMDATA, &items)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ITEMDATA, err.Error())
		return &Item{}, errors.New(s)
	}

	i := items[strings.ToLower(name)]

	return &i, nil
}

func RegisterTrainerTalent(name string, classSpecific bool, requirements, frequency, target, description string, continuous, standart, free, interrupt, extended, legal bool) (*TrainerTalent, error) {
	talent := &TrainerTalent{
		Name: strings.Title(name),

		IsClassSpecific: classSpecific,

		Requirements: requirements,
		Frequency:    frequency,
		Target:       target,
		Description:  description,

		IsContinuous: continuous,
		IsStandart:   standart,
		IsFree:       free,
		IsInterrupt:  interrupt,
		IsExtended:   extended,
		IsLegal:      legal,
	}

	var talents map[string]TrainerTalent
	err := general.GetJsonData(TALENTDATA, &talents)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", TALENTDATA, err.Error())
		return nil, errors.New(s)
	}

	talents[strings.ToLower(name)] = *talent

	err = general.SetJsonData(TALENTDATA, talents)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", TALENTDATA, err.Error())
		return nil, errors.New(s)
	}

	return talent, nil
}

func GetTrainerTalent(name string) (*TrainerTalent, error) {
	var talents map[string]TrainerTalent
	err := general.GetJsonData(TALENTDATA, &talents)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", TALENTDATA, err.Error())
		return &TrainerTalent{}, errors.New(s)
	}

	TT := talents[strings.ToLower(name)]

	return &TT, nil
}

func RegisterTrainerClass(name, description, parentClass string, basicTalents [2]*TrainerTalent, possibleTalents []*TrainerTalent, expertise []*Expertise, requirements string) (*TrainerClass, error) {
	newClass := &TrainerClass{
		Name:        strings.Title(name),
		Description: description,
		ParentClass: parentClass,

		BasicTalents:    basicTalents,
		PossibleTalents: possibleTalents,

		Expertises:   expertise,
		Requirements: requirements,
	}

	var classes map[string]TrainerClass
	err := general.GetJsonData(CLASSDATA, &classes)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CLASSDATA, err.Error())
		return nil, errors.New(s)
	}

	classes[strings.ToLower(name)] = *newClass

	err = general.SetJsonData(CLASSDATA, classes)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", CLASSDATA, err.Error())
		return nil, errors.New(s)
	}

	return newClass, nil
}

func GetTrainerClass(name string) (*TrainerClass, error) {
	var classes map[string]TrainerClass
	err := general.GetJsonData(CLASSDATA, &classes)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CLASSDATA, err.Error())
		return &TrainerClass{}, errors.New(s)
	}

	TC := classes[strings.ToLower(name)]

	return &TC, nil
}

func RegisterExpertise(name string, associatedStat, description string) (*Expertise, error) {
	newExpertise := &Expertise{
		Name:        strings.Title(name),
		Description: description,

		Double: false,

		AssociatedStat: associatedStat,
	}

	var expertises map[string]Expertise
	err := general.GetJsonData(EXPERTISEDATA, &expertises)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", EXPERTISEDATA, err.Error())
		return nil, errors.New(s)
	}

	expertises[strings.ToLower(name)] = *newExpertise

	err = general.SetJsonData(EXPERTISEDATA, expertises)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", EXPERTISEDATA, err.Error())
		return nil, errors.New(s)
	}

	return newExpertise, nil
}

func GetExpertise(name string) (*Expertise, error) {
	var expertises map[string]Expertise
	err := general.GetJsonData(EXPERTISEDATA, &expertises)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", EXPERTISEDATA, err.Error())
		return &Expertise{}, errors.New(s)
	}

	exp := expertises[strings.ToLower(name)]

	return &exp, nil
}

func RegisterCapacity(name, description string) (*Capacity, error) {
	newCapacity := &Capacity{
		Name:        strings.Title(name),
		Description: description,
	}

	var capacities map[string]Capacity
	err := general.GetJsonData(CAPACITYDATA, &capacities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CAPACITYDATA, err.Error())
		return nil, errors.New(s)
	}

	capacities[strings.ToLower(name)] = *newCapacity

	err = general.SetJsonData(CAPACITYDATA, capacities)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", CAPACITYDATA, err.Error())
		return nil, errors.New(s)
	}

	return newCapacity, nil
}

func GetCapacity(name string) (*Capacity, error) {
	var capacities map[string]Capacity
	err := general.GetJsonData(CAPACITYDATA, &capacities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CAPACITYDATA, err.Error())
		return &Capacity{}, errors.New(s)
	}

	cap := capacities[strings.ToLower(name)]

	return &cap, nil
}

////////////////////////////////////////////////////////////////////////

func newPokemonStatusTable(stats map[string]int) (*PokemonStatusTable, error) {
	{
		if _, ok := stats["HP"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}

		if _, ok := stats["ATK"]; !ok {
			return nil, errors.New("Invalid stats (missing ATK)\n")
		}

		if _, ok := stats["DEF"]; !ok {
			return nil, errors.New("Invalid stats (missing DEF)\n")
		}

		if _, ok := stats["SPATK"]; !ok {
			return nil, errors.New("Invalid stats (missing SPATK)\n")
		}

		if _, ok := stats["SPDEF"]; !ok {
			return nil, errors.New("Invalid stats (missing SPDEF)\n")
		}

		if _, ok := stats["SPD"]; !ok {
			return nil, errors.New("Invalid stats (missing SPD)\n")
		}
	}

	zero := map[string]int{"HP": 0, "ATK": 0, "DEF": 0, "SPATK": 0, "SPDEF": 0, "SPD": 0}

	var keys [6]string = [6]string{"HP", "ATK", "DEF", "SPATK", "SPDEF", "SPD"}

	sort.SliceStable(keys[:], func(i, j int) bool {
		return (stats[keys[i]] > stats[keys[j]])
	})

	table := &PokemonStatusTable{
		Base:   stats,
		LvlUp:  zero,
		Total:  stats,
		Stages: zero,

		BaseRelation: keys,

		Distributable: [2]int{0, 0},
	}

	return table, nil
}

func newTrainerStatusTable(stats map[string]int, lvl int) (*TrainerStatusTable, error) {
	{
		if _, ok := stats["HP"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}

		if _, ok := stats["ATK"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}

		if _, ok := stats["DEF"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}

		if _, ok := stats["SPATK"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}

		if _, ok := stats["SPDEF"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}

		if _, ok := stats["SPD"]; !ok {
			return nil, errors.New("Invalid stats (missing HP)\n")
		}
	}

	zero := map[string]int{"HP": 0, "ATK": 0, "DEF": 0, "SPATK": 0, "SPDEF": 0, "SPD": 0}

	var keys [6]string = [6]string{"HP", "ATK", "DEF", "SPATK", "SPDEF", "SPD"}
	modifiers := map[string]int{"HP": 0, "ATK": 0, "DEF": 0, "SPATK": 0, "SPDEF": 0, "SPD": 0}

	for _, key := range keys {
		if stats[key] < 10 {
			modifiers[key] = stats[key] - 10
		} else {
			modifiers[key] = (stats[key] - 10) / 2
		}
	}

  modifiers["HP"] = 0

	statTotal := stats["HP"] + stats["ATK"] + stats["DEF"] + stats["SPATK"] + stats["SPDEF"] + stats["SPD"]

	table := &TrainerStatusTable{
		Status:    stats,
		Modifiers: modifiers,
		Stages:    zero,
		Total:     stats,

		Distributable: [2]int{statTotal, 66 + TRAINERLVLTABLE["total_status"][lvl]},
	}

	return table, nil
}

func newCapacityTable(capacities [3]int, others []*Capacity) *CapacityTable {
	table := &CapacityTable{
		Strength:    capacities[0],
		Inteligence: capacities[1],
		Jump:        capacities[2],

		Others: others,
	}

	return table
}

////////////////////////////////////////////////////////////////////////

func CreatePokemonSheet(nickname, species, gender, nature string, abilities []*PokemonAbility, lvl int) (*PokemonSheet, error) {
	speciesData, err := GetSpecies(species)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", SPECIESDATA, err.Error())
		return nil, errors.New(s)
	}

	if speciesData.Number == 0 {
		return nil, errors.New("Species not registered: number = 0")
	}

	randomFactor := (&general.DiceSet{X: 10, N: 20, Mod: -10}).Roll()
	H := int((float32(randomFactor)/float32(190)*0.4-0.2)*float32(speciesData.AverageHeight)) + speciesData.AverageHeight

	randomFactor = (&general.DiceSet{X: 10, N: 20, Mod: -10}).Roll()
	W := int((float32(randomFactor)/float32(190)*0.4-0.2)*float32(speciesData.AverageWeight)) + speciesData.AverageWeight

	status, err := newPokemonStatusTable(speciesData.BaseStats)
	if err != nil {
		return nil, err
	}

	status.Distributable = [2]int{0, lvl - 1}

	aux, err := general.GetRD("sheetCount")
	if aux == "" {
		general.SetRD("sheetCount", "0")
	} else if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(aux)
	if err != nil {
		return nil, err
	}

	newSheet := &PokemonSheet{
		Id: id,

		Nick:    nickname,
		Species: speciesData,
		Height:  H,
		Weight:  W,
		Gender:  gender,

		Nature: nature,

		Lvl: lvl,
		Exp: 0,

		Status: status,
		Hp:     [2]int{(lvl + status.Base["HP"]) * 4},

		Movement:  speciesData.Movement,
		Evasion:   [3]int{general.Capped(status.Base["DEF"]/5, 0, 6), general.Capped(status.Base["SPDEF"]/5, 0, 6), general.Capped(status.Base["SPD"]/5, 0, 6)},
		ElemBonus: lvl / 5,

		Abilities: abilities,
		Moves:     [2][4]*PokemonMove{},

		Notes: "",
	}
  newSheet.Hp[1] = newSheet.Hp[0]

	err = general.SetJsonData("./data/sheets/"+aux+"_1.json", newSheet)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", "./data/sheets/"+aux+"_1.json", err.Error())
		return nil, errors.New(s)
	}

	general.SetRD("sheetCount", strconv.Itoa(id+1))

	return newSheet, nil
}

func CreateTrainerSheet(name, player, gender string, lvl, age, height, weight int, stats map[string]int) (*TrainerSheet, error) {
	stts, err := newTrainerStatusTable(stats, lvl)
	if err != nil {
		return nil, err
	}

	movement := map[string]int{"land": 0, "swimming": general.Capped(stts.Modifiers["DEF"]/2+3, 4, 100), "underwater": 0}

	pokedex, err := GetTrainerTalent("Pokéagenda")
	if err != nil {
		return nil, err
	}

	weapons, err := GetTrainerTalent("armas")
	if err != nil {
		return nil, err
	}

	stdTalents := []*TrainerTalent{pokedex, weapons}

	if stts.Modifiers["ATK"] >= stts.Modifiers["SPD"] {
		movement["land"] = general.Capped(stts.Modifiers["ATK"]/2+2, 4, 100)
	} else {
		movement["land"] = general.Capped(stts.Modifiers["SPD"]/2+2, 4, 100)
	}

	if stts.Modifiers["ATK"] >= 3 || stts.Modifiers["DEF"] >= 3 {
		movement["underwater"] = 4
	} else {
		movement["underwater"] = 3
	}

	aux, err := general.GetRD("sheetCount")
	if aux == "" {
		general.SetRD("sheetCount", "0")
	} else if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(aux)
	if err != nil {
		return nil, err
	}

  talentSlots := TRAINERLVLTABLE["total_talents"][lvl]

	newSheet := &TrainerSheet{
		Id: id,

		Name:   name,
		Player: player,

		Gender: gender,
		Age:    age,
		Height: height,
		Weight: weight,

		Lvl: lvl,

		Status: stts,
		Hp:     [2]int{(lvl + stats["HP"]) * 4},

		Movement:             movement,
		Evasion:              [3]int{general.Capped(stts.Status["DEF"]/5, 0, 6), general.Capped(stts.Status["SPDEF"]/5, 0, 6), general.Capped(stts.Status["SPD"]/5, 0, 6)},
		WeaponDamageCategory: lvl/7 + 1,
		WeaponDamage:         WEAPONDAMAGETABLE[lvl/7],

		Talents:     stdTalents,
		TalentSlots: talentSlots,
    Expertises: []Expertise{},
	}
  newSheet.Hp[1] = newSheet.Hp[0]

	err = general.SetRD("sheetCount", strconv.Itoa(id+1))
	if err != nil {
		return nil, err
	}

	err = general.SetJsonData(general.SHEETSPATH+aux+"_0.json", newSheet)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", general.SHEETSPATH+aux+"_0.json", err.Error())
		return nil, errors.New(s)
	}

	return newSheet, nil
}
