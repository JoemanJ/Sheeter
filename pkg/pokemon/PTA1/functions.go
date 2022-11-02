package PTA1

import (
	"sort"
	"strings"

	"github.com/mtslzr/pokeapi-go"

	actions "Joe/sheet-hole/pkg/general"
)

const ABILITYDATA string = "./data/testeAbilities.json"
const MOVEDATA string = "./data/testeMoves.json"
const SPECIESDATA string = "./data/speciesData.json"
const ITEMDATA string = "./data/itemData.json"

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

func RegisterSpecies(name string, diet string, abilities []*PokemonAbility) (*PokemonSpecies, error) {
	pokemon, err := pokeapi.Pokemon(strings.ToLower(name))
	if err != nil {
		return nil, err
	}

	typesData := pokemon.Types

	var types []string

	for _, t := range typesData {
		types = append(types, t.Type.Name)
	}

	newSpecies := &PokemonSpecies{
		Number: pokemon.ID,
		Name:   strings.Title(pokemon.Species.Name),

		Type: types,
		Diet: diet,

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

func RegisterItem(name string, description string) *Item {
	i := &Item{
		Quantity: 0,

		Name:        name,
		Description: description,
	}

	var items map[string]Item

	getJsonData(ITEMDATA, &items)

	items[strings.ToLower(name)] = *i

	setJsonData(ITEMDATA, items)

	return i
}

func GetItem(name string) (Item, error) {
	var items map[string]Item
	err := getJsonData(ITEMDATA, &items)
	if err != nil {
		return Item{}, err
	}

	return items[strings.ToLower(name)], nil
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
	var modifiers map[string]int

	for _, key := range keys {
		if stats[key] < 10 {
			modifiers[key] = stats[key] - 10
		} else {
			modifiers[key] = (stats[key] - 10) / 2
		}
	}

	table := &TrainerStatusTable{
		Status:    stats,
		Modifiers: modifiers,
		Stages:    zero,
		Total:     stats,

		Distributable: [2]int{0, 66},
	}

	return table
}
