package PTA1

import (
	actions "Joe/sheet-hole/pkg/general"
	"fmt"
	"html/template"
	"log"
)

var WEAPONDAMAGETABLE [8]*actions.DiceSet = [8]*actions.DiceSet{{X: 1, N: 10, Mod: 4}, {X: 1, N: 12, Mod: 6}, {X: 2, N: 8, Mod: 6}, {X: 2, N: 10, Mod: 8}, {X: 3, N: 8, Mod: 10}, {X: 3, N: 10, Mod: 12}, {X: 3, N: 12, Mod: 14}, {X: 4, N: 12, Mod: 16}}

var TRAINERLVLTABLE map[string][51]int = map[string][51]int{
	"classes": {0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	"status":  {0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 2, 0, 1, 0, 3},
	"talents": {0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
}

// POKEMON STRUCTURES
type PokemonSpecies struct {
	Number int
	Name   string

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
	Damage      *actions.DiceSet
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
	WeaponDamage         *actions.DiceSet

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

func (s *TrainerSheet) SheetBody() (*template.Template, error) {
	t, err := template.New("sheet_body").ParseFiles("./pkg/pokemon/PTA1/trainerSheet.partial.html")

	log.Printf("\n\n%v\n\n", t)

	if err != nil {
		return template.New("sheet_body").Parse(fmt.Sprintf("<div>Could not open sheet!\nerror: %s</div>", err.Error()))
	}

	return t, nil
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

	Others []Capacity
}
