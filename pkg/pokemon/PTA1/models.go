package PTA1

import (
	general "Joe/sheeter/pkg/general"
	sheeters "Joe/sheeter/pkg/general"
	"bytes"
	"html/template"
	"net/http"
)

const TRAINER_SHEETID = 0
const POKEMON_SHEETID = 1

var WEAPONDAMAGETABLE [8]*general.DiceSet = [8]*general.DiceSet{{X: 1, N: 10, Mod: 4}, {X: 1, N: 12, Mod: 6}, {X: 2, N: 8, Mod: 6}, {X: 2, N: 10, Mod: 8}, {X: 3, N: 8, Mod: 10}, {X: 3, N: 10, Mod: 12}, {X: 3, N: 12, Mod: 14}, {X: 4, N: 12, Mod: 16}}

var TRAINERLVLTABLE map[string][51]int = map[string][51]int{
	"classes":       {0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	"status":        {0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 3},
	"talents":       {0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
	"total_status":  {0, 1, 2, 3, 4, 5, 6, 6, 7, 7, 9, 9, 10, 10, 12, 12, 13, 13, 15, 15, 16, 16, 18, 18, 19, 19, 21, 21, 22, 22, 24, 24, 25, 25, 27, 27, 28, 28, 30, 30, 31, 31, 33, 33, 34, 34, 36, 36, 37, 37, 39},
	"total_talents": {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11, 12, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31},
}

// , 10POKEMON STRUCTURES
type PokemonSpecies struct {
	Number int
	Name   string

	Sprite string

	Type       []string
	Diet       string
	Capacities *CapacityTable
	Movement   map[string]int

	AverageHeight int
	AverageWeight int
	BaseStats     map[string]int

	Abilities     []*PokemonAbility
	HighAbilities []*PokemonAbility
}

type PokemonStatusTable struct {
	Base   map[string]int
	LvlUp  map[string]int
	Total  map[string]int
	Stages map[string]int

	BaseRelation [6]string

	Distributable [2]int
}

type PokemonMove struct {
	Name        string
	Type        string
	Aptitude    string
	Descriptors []string
	AccDiff     int
	Damage      *general.DiceSet
	Reach       string
	Frequency   string
	Contests    string
	Effect      string
}

type PokemonAbility struct {
	Name        string
	Activation  string
	Description string
}

type PokemonSheet struct {
	Id int

	Nick    string
	Species string
	Height  int
	Weight  int
	Gender  string

  Sprite string

  Type []string

	Nature string

	Lvl int
	Exp int

	Status *PokemonStatusTable
	Hp     [2]int

	Movement  map[string]int
	Evasion   [3]int
	ElemBonus int

	Abilities []*PokemonAbility
	Moves     [2][4]*PokemonMove

	Notes string
}

func (s *PokemonSheet) Render(w http.ResponseWriter) error {
	template := new(template.Template)

	template, err := sheeters.RenderVolatile("sheet.page.html", "./ui/html")
	if err != nil {
		return err
	}

	template, err = template.ParseFiles("./pkg/pokemon/PTA1/html/pokemonSheet.partial.html")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	err = template.Execute(buf, nil)
	if err != nil {
		return err
	}

	buf.WriteTo(w)

	return nil
}

// TRAINER STRUCTURES
type TrainerSheet struct {
	Id int

	Name   string
	Player string

	Gender string
	Age    int
	Height int
	Weight int

	Lvl     int
	Classes [4]*TrainerClass

	Status *TrainerStatusTable
	Hp     [2]int

	Movement             map[string]int
	Evasion              [3]int
	WeaponDamageCategory int
	WeaponDamage         *general.DiceSet

	Talents     []*TrainerTalent
	TalentSlots int

	TotalSeenPokemon   int
	SeenPokemon        [807]bool
	TotalCaughtPokemon int
	CaughtPokemon      [807]bool

	Inventory []*Item

	PokemonList []*PokemonSheet

	Prizes string

	Notes string
}

func (s *TrainerSheet) Render(w http.ResponseWriter) error {
	tmpl, err := sheeters.RenderVolatile("sheet.page.html", "./ui/html")
	if err != nil {
		return err
	}

	tmpl, err = tmpl.ParseFiles("./pkg/pokemon/PTA1/html/trainerSheet.partial.html")

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, s)
	if err != nil {
		return err
	}

	buf.WriteTo(w)

	return nil
}

type TrainerClass struct {
	Name        string
	Description string

	ParentClass string

	BasicTalents    [2]*TrainerTalent
	PossibleTalents []*TrainerTalent

	Expertises   []*Expertise
	Requirements string
}

type TrainerStatusTable struct {
	Status    map[string]int
	Modifiers map[string]int
	Total     map[string]int
	Stages    map[string]int

	Distributable [2]int
}

type TrainerTalent struct {
	Name string

	IsClassSpecific bool

	Requirements string

	Frqeuency string
	Target    string

	Description string

	IsContinuous bool
	IsStandart   bool
	IsFree       bool
	IsInterrupt  bool
	IsExtended   bool
	IsLegal      bool
}

type Item struct {
	Quantity int

	Name        string
	Description string
}

type Expertise struct {
	Double bool

	AssociatedStat string

	Name        string
	Description string
}

type Capacity struct {
	Name        string
	Description string
}

type CapacityTable struct {
	Strength    int
	Inteligence int
	Jump        int

	Others []*Capacity
}
