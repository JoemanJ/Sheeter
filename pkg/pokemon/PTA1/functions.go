package PTA1

import (
	"sort"
	"strings"

	"github.com/mtslzr/pokeapi-go"

	actions "Joe/sheet-hole/pkg/general"
)

const ABILITYDATA string = "./data/abilitieData.json"
const MOVEDATA string = "./data/moveData.json"
const SPECIESDATA string = "./data/speciesData.json"
const ITEMDATA string = "./data/itemData.json"
const TALENTDATA string = "./data/talentData.json"
const CLASSDATA string = "./data/classData.json"
const EXPERTISEDATA string = "./data/expertiseData.json"
const CAPACITYDATA string = "./data/capacityData.json"

func RegisterAbility(name string, activation string, description string) (*PokemonAbility, error) {
	newAbility := &PokemonAbility{
		IsHighAbility: false,
		Name:          strings.Title(name),
		Activation:    activation,
		Description:   description,
	}

	var abilities map[string]PokemonAbility

	err := getJsonData(ABILITYDATA, &abilities)
	if err != nil {
		return nil, err
	}

	abilities[strings.ToLower(name)] = *newAbility

	setJsonData(ABILITYDATA, abilities)

	return newAbility, nil
}

func GetAbility(name string) (PokemonAbility, error) {
	var abilities map[string]PokemonAbility
	err := getJsonData(ABILITYDATA, &abilities)
	if err != nil {
		return PokemonAbility{}, err
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
		return nil, err
	}

	moves[strings.ToLower(name)] = *newMove

	err = setJsonData(MOVEDATA, moves)
	if err != nil {
		return nil, err
	}

	return newMove, nil
}

func GetMove(name string) (PokemonMove, error) {
	var moves map[string]PokemonMove
	err := getJsonData(MOVEDATA, &moves)
	if err != nil {
		return PokemonMove{}, err
	}

	return moves[strings.ToLower(name)], nil
}

func RegisterSpecies(name string, diet string, capacities [3]int, others []Capacity, abilities []*PokemonAbility) (*PokemonSpecies, error) {
	pokemon, err := pokeapi.Pokemon(strings.ToLower(name))
	if err != nil {
		return nil, err
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

		AverageHeight: pokemon.Height,
		AverageWeight: pokemon.Weight,
		BaseStats:     map[string]int{"HP": (pokemon.Stats[0].BaseStat + 5) / 10, "ATK": (pokemon.Stats[1].BaseStat + 5) / 10, "DEF": (pokemon.Stats[2].BaseStat + 5) / 10, "SPATK": (pokemon.Stats[3].BaseStat + 5) / 10, "SPDEF": (pokemon.Stats[4].BaseStat + 5) / 10, "SPD": (pokemon.Stats[5].BaseStat + 5) / 10},

		Abilities: abilities,
	}

	var species map[string]PokemonSpecies

	err = getJsonData(SPECIESDATA, &species)
	if err != nil {
		return nil, err
	}

	species[strings.ToLower(name)] = *newSpecies

	err = setJsonData(SPECIESDATA, species)

	return newSpecies, nil
}

func GetSpecies(name string) (PokemonSpecies, error) {
	var species map[string]PokemonSpecies
	err := getJsonData(SPECIESDATA, &species)
	if err != nil {
		return PokemonSpecies{}, err
	}

	return species[strings.ToLower(name)], nil
}

func RegisterItem(name string, description string) (*Item, error) {
	i := &Item{
		Quantity: 0,

		Name:        name,
		Description: description,
	}

	var items map[string]Item

	err := getJsonData(ITEMDATA, &items)
	if err != nil {
		return nil, err
	}

	items[strings.ToLower(name)] = *i

	err = setJsonData(ITEMDATA, items)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func GetItem(name string) (Item, error) {
	var items map[string]Item
	err := getJsonData(ITEMDATA, &items)
	if err != nil {
		return Item{}, err
	}

	return items[strings.ToLower(name)], nil
}

func RegisterTrainerTalent(name string, classSpecific bool, requirements, frequency, target, description string, continuous, standart, free, interrupt, extended, legal bool) (*TrainerTalent, error) {
	talent := &TrainerTalent{
		Name: name,

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
		return nil, err
	}

	talents[strings.ToLower(name)] = *talent

	err = setJsonData(TALENTDATA, talents)
	if err != nil {
		return nil, err
	}

	return talent, nil
}

func GetTrainerTalent(name string) (TrainerTalent, error) {
	var talents map[string]TrainerTalent
	err := getJsonData(TALENTDATA, *&talents)
	if err != nil {
		return TrainerTalent{}, err
	}

	return talents[strings.ToLower(name)], nil
}

func RegisterTrainerClass(name, description, parentClass string, basicTalents [2]*TrainerTalent, possibleTalents []*TrainerTalent, expertise []*Expertise, requirements string) (*TrainerClass, error) {
	newClass := &TrainerClass{
		Name:        name,
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
		return nil, err
	}

	classes[strings.ToLower(name)] = *newClass

	err = setJsonData(CLASSDATA, classes)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}

func GetTrainerClass(name string) (TrainerClass, error) {
	var classes map[string]TrainerClass
	err := getJsonData(CLASSDATA, &classes)
	if err != nil {
		return TrainerClass{}, err
	}

	return classes[strings.ToLower(name)], nil
}

func RegisterExpertise(name string, associatedStat, description string) (*Expertise, error) {
	newExpertise := &Expertise{
		Name:        name,
		Description: description,

		Double: false,

		AssociatedStat: associatedStat,
	}

	var expertises map[string]Expertise
	err := getJsonData(EXPERTISEDATA, &expertises)
	if err != nil {
		return nil, err
	}

	expertises[strings.ToLower(name)] = *newExpertise

	err = setJsonData(EXPERTISEDATA, expertises)
	if err != nil {
		return nil, err
	}

	return newExpertise, nil
}

func GetExpertise(name string) (Expertise, error) {
	var expertises map[string]Expertise
	err := getJsonData(EXPERTISEDATA, &expertises)
	if err != nil {
		return Expertise{}, err
	}

	return expertises[strings.ToLower(name)], nil
}

func RegisterCapacity(name, description string) (*Capacity, error) {
	newCapacity := &Capacity{
		Name:        name,
		Description: description,
	}

	var capacities map[string]Capacity
	err := getJsonData(CAPACITYDATA, &capacities)
	if err != nil {
		return nil, err
	}

	capacities[strings.ToLower(name)] = *newCapacity

	err = setJsonData(CAPACITYDATA, capacities)
	if err != nil {
		return nil, err
	}

	return newCapacity, nil
}

func GetCapacity(name string) (Capacity, error) {
	var capacities map[string]Capacity
	err := getJsonData(CAPACITYDATA, &capacities)
	if err != nil {
		return Capacity{}, err
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
