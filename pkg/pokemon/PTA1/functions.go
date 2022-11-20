package PTA1

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/mtslzr/pokeapi-go"

	actions "Joe/sheet-hole/pkg/general"
)

const ABILITYDATA string = "./data/abilityData.json"
const MOVEDATA string = "./data/moveData.json"
const SPECIESDATA string = "./data/speciesData.json"
const ITEMDATA string = "./data/itemData.json"
const TALENTDATA string = "./data/talentData.json"
const CLASSDATA string = "./data/classData.json"
const EXPERTISEDATA string = "./data/expertiseData.json"
const CAPACITYDATA string = "./data/capacityData.json"

func RegisterAbility(name string, activation string, description string) (*PokemonAbility, error) {
	newAbility := &PokemonAbility{
		Name:        strings.Title(name),
		Activation:  activation,
		Description: description,
	}

	var abilities map[string]PokemonAbility

	err := getJsonData(ABILITYDATA, &abilities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ABILITYDATA, err.Error())
		return nil, errors.New(s)
	}

	abilities[strings.ToLower(name)] = *newAbility

	err = setJsonData(ABILITYDATA, abilities)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", ABILITYDATA, err.Error())
		return nil, errors.New(s)
	}

	return newAbility, nil
}

func GetAbility(name string) (PokemonAbility, error) {
	var abilities map[string]PokemonAbility
	err := getJsonData(ABILITYDATA, &abilities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ABILITYDATA, err.Error())
		return PokemonAbility{}, errors.New(s)
	}

	return abilities[strings.ToLower(name)], nil
}

func RegisterMove(name string, Type string, aptitude string, descriptors []string, accDiff int, dice *actions.DiceSet, reach string, frequency string, contests string, effect string) (*PokemonMove, error) {
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

	err := getJsonData(MOVEDATA, &moves)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", MOVEDATA, err.Error())
		return nil, errors.New(s)
	}

	moves[strings.ToLower(name)] = *newMove

	err = setJsonData(MOVEDATA, moves)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", MOVEDATA, err.Error())
		return nil, errors.New(s)
	}

	return newMove, nil
}

func GetMove(name string) (PokemonMove, error) {
	var moves map[string]PokemonMove
	err := getJsonData(MOVEDATA, &moves)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", MOVEDATA, err.Error())
		return PokemonMove{}, errors.New(s)
	}

	return moves[strings.ToLower(name)], nil
}

func RegisterSpecies(name string, diet string, capacities [3]int, others []Capacity, abilities []*PokemonAbility, highAbilities []*PokemonAbility, movement map[string]int) (*PokemonSpecies, error) {
	pokemon, err := pokeapi.Pokemon(strings.ToLower(name))
	if err != nil {
		s := fmt.Sprintf("Error getting data from PokeAPI:\n%s", err.Error())
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

	err = getJsonData(SPECIESDATA, &species)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", SPECIESDATA, err.Error())
		return nil, errors.New(s)
	}

	species[strings.ToLower(name)] = *newSpecies

	err = setJsonData(SPECIESDATA, species)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", SPECIESDATA, err.Error())
		return nil, errors.New(s)
	}

	return newSpecies, nil
}

func GetSpecies(name string) (PokemonSpecies, error) {
	var species map[string]PokemonSpecies
	err := getJsonData(SPECIESDATA, &species)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", SPECIESDATA, err.Error())
		return PokemonSpecies{}, errors.New(s)
	}

	return species[strings.ToLower(name)], nil
}

func RegisterItem(name string, description string) (*Item, error) {
	i := &Item{
		Quantity: 0,

		Name:        strings.Title(name),
		Description: description,
	}

	var items map[string]Item

	err := getJsonData(ITEMDATA, &items)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ITEMDATA, err.Error())
		return nil, errors.New(s)
	}

	items[strings.ToLower(name)] = *i

	err = setJsonData(ITEMDATA, items)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", ITEMDATA, err.Error())
		return nil, errors.New(s)
	}

	return i, nil
}

func GetItem(name string) (Item, error) {
	var items map[string]Item
	err := getJsonData(ITEMDATA, &items)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", ITEMDATA, err.Error())
		return Item{}, errors.New(s)
	}

	return items[strings.ToLower(name)], nil
}

func RegisterTrainerTalent(name string, classSpecific bool, requirements, frequency, target, description string, continuous, standart, free, interrupt, extended, legal bool) (*TrainerTalent, error) {
	talent := &TrainerTalent{
		Name: strings.Title(name),

		IsClassSpecific: classSpecific,

		Requirements: requirements,
		Frqeuency:    frequency,
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
	err := getJsonData(TALENTDATA, &talents)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", TALENTDATA, err.Error())
		return nil, errors.New(s)
	}

	talents[strings.ToLower(name)] = *talent

	err = setJsonData(TALENTDATA, talents)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", TALENTDATA, err.Error())
		return nil, errors.New(s)
	}

	return talent, nil
}

func GetTrainerTalent(name string) (TrainerTalent, error) {
	var talents map[string]TrainerTalent
	err := getJsonData(TALENTDATA, *&talents)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", TALENTDATA, err.Error())
		return TrainerTalent{}, errors.New(s)
	}

	return talents[strings.ToLower(name)], nil
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
	err := getJsonData(CLASSDATA, &classes)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CLASSDATA, err.Error())
		return nil, errors.New(s)
	}

	classes[strings.ToLower(name)] = *newClass

	err = setJsonData(CLASSDATA, classes)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", CLASSDATA, err.Error())
		return nil, errors.New(s)
	}

	return newClass, nil
}

func GetTrainerClass(name string) (TrainerClass, error) {
	var classes map[string]TrainerClass
	err := getJsonData(CLASSDATA, &classes)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CLASSDATA, err.Error())
		return TrainerClass{}, errors.New(s)
	}

	return classes[strings.ToLower(name)], nil
}

func RegisterExpertise(name string, associatedStat, description string) (*Expertise, error) {
	newExpertise := &Expertise{
		Name:        strings.Title(name),
		Description: description,

		Double: false,

		AssociatedStat: associatedStat,
	}

	var expertises map[string]Expertise
	err := getJsonData(EXPERTISEDATA, &expertises)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", EXPERTISEDATA, err.Error())
		return nil, errors.New(s)
	}

	expertises[strings.ToLower(name)] = *newExpertise

	err = setJsonData(EXPERTISEDATA, expertises)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", EXPERTISEDATA, err.Error())
		return nil, errors.New(s)
	}

	return newExpertise, nil
}

func GetExpertise(name string) (Expertise, error) {
	var expertises map[string]Expertise
	err := getJsonData(EXPERTISEDATA, &expertises)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", EXPERTISEDATA, err.Error())
		return Expertise{}, errors.New(s)
	}

	return expertises[strings.ToLower(name)], nil
}

func RegisterCapacity(name, description string) (*Capacity, error) {
	newCapacity := &Capacity{
		Name:        strings.Title(name),
		Description: description,
	}

	var capacities map[string]Capacity
	err := getJsonData(CAPACITYDATA, &capacities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CAPACITYDATA, err.Error())
		return nil, errors.New(s)
	}

	capacities[strings.ToLower(name)] = *newCapacity

	err = setJsonData(CAPACITYDATA, capacities)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", CAPACITYDATA, err.Error())
		return nil, errors.New(s)
	}

	return newCapacity, nil
}

func GetCapacity(name string) (Capacity, error) {
	var capacities map[string]Capacity
	err := getJsonData(CAPACITYDATA, &capacities)
	if err != nil {
		s := fmt.Sprintf("Error reading file %s:\n%s", CAPACITYDATA, err.Error())
		return Capacity{}, errors.New(s)
	}

	return capacities[strings.ToLower(name)], nil
}

////////////////////////////////////////////////////////////////////////

func newPokemonStatusTable(stats map[string]int) *PokemonStatusTable {
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

	return table
}

func newTrainerStatusTable(stats map[string]int) *TrainerStatusTable {
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

	statTotal := stats["HP"] + stats["ATK"] + stats["DEF"] + stats["SPATK"] + stats["SPDEF"] + stats["SPD"]

	table := &TrainerStatusTable{
		Status:    stats,
		Modifiers: modifiers,
		Stages:    zero,
		Total:     stats,

		Distributable: [2]int{statTotal, 66},
	}

	return table
}

func newCapacityTable(capacities [3]int, others []Capacity) *CapacityTable {
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

	randomFactor := actions.RollSet(&actions.DiceSet{X: 10, N: 20, Mod: -10})
	H := int((float32(randomFactor)/float32(190)*0.4-0.2)*float32(speciesData.AverageHeight)) + speciesData.AverageHeight

	randomFactor = actions.RollSet(&actions.DiceSet{X: 10, N: 20, Mod: -10})
	W := int((float32(randomFactor)/float32(190)*0.4-0.2)*float32(speciesData.AverageWeight)) + speciesData.AverageWeight

	status := newPokemonStatusTable(speciesData.BaseStats)
	status.Distributable = [2]int{0, lvl - 1}

	newSheet := &PokemonSheet{
		Nick:    nickname,
		Species: species,
		Height:  H,
		Weight:  W,
		Gender:  gender,

		Nature: nature,

		Lvl: lvl,
		Exp: 0,

		Status: status,
		Hp:     [2]int{(lvl + status.Base["HP"]) * 4},

		Movement:  speciesData.Movement,
		Evasion:   [3]int{actions.Capped(status.Base["DEF"]/5, 0, 6), actions.Capped(status.Base["SPDEF"]/5, 0, 6), actions.Capped(status.Base["SPD"]/5, 0, 6)},
		ElemBonus: lvl / 5,

		Abilities: abilities,
		Moves:     [2][4]*PokemonMove{},

		Notes: "",
	}

	err = setJsonData("./data/sheets/"+nickname+"_sheet.json", newSheet)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", "./data/sheets/"+nickname+"_sheet.json", err.Error())
		return nil, errors.New(s)
	}

	return newSheet, nil
}

func CreateTrainerSheet(name, player, gender string, lvl, age, height, weight int, stats map[string]int) (*TrainerSheet, error) {
	stts := newTrainerStatusTable(stats)
	movement := map[string]int{"land": 0, "swimming": actions.Capped(stts.Modifiers["DEF"]/2+3, 4, 100), "underwater": 0}

	pokedex, err := GetTrainerTalent("Pokéagenda")
	if err != nil {
		return nil, err
	}

	weapons, err := GetTrainerTalent("armas")
	if err != nil {
		return nil, err
	}

	stdTalents := []*TrainerTalent{&pokedex, &weapons}

	if stts.Modifiers["ATK"] >= stts.Modifiers["SPD"] {
		movement["land"] = actions.Capped(stts.Modifiers["ATK"]/2+2, 4, 100)
	} else {
		movement["land"] = actions.Capped(stts.Modifiers["SPD"]/2+2, 4, 100)
	}

	if stts.Modifiers["ATK"] >= 3 || stts.Modifiers["DEF"] >= 3 {
		movement["underwater"] = 4
	} else {
		movement["underwater"] = 3
	}

	newSheet := &TrainerSheet{
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
		Evasion:              [3]int{actions.Capped(stts.Status["DEF"]/5, 0, 6), actions.Capped(stts.Status["SPDEF"]/5, 0, 6), actions.Capped(stts.Status["SPD"]/5, 0, 6)},
		WeaponDamageCategory: lvl/7 + 1,
		WeaponDamage:         WEAPONDAMAGETABLE[lvl/7],

		Talents:     stdTalents,
		TalentSlots: 0,
	}

	return newSheet, nil
}
