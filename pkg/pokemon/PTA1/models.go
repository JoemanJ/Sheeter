package PTA1

import (
	general "Joe/sheeter/pkg/general"
	sheeters "Joe/sheeter/pkg/general"
	"bytes"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

const TRAINER_SHEETID = 0
const POKEMON_SHEETID = 1

var WEAPONDAMAGETABLE [8]*general.DiceSet = [8]*general.DiceSet{{X: 1, N: 10, Mod: 4}, {X: 1, N: 12, Mod: 6}, {X: 2, N: 8, Mod: 6}, {X: 2, N: 10, Mod: 8}, {X: 3, N: 8, Mod: 10}, {X: 3, N: 10, Mod: 12}, {X: 3, N: 12, Mod: 14}, {X: 4, N: 12, Mod: 16}}
var POKEMONEXPTABLE = [100]int{0, 25, 50, 100, 150, 200, 400, 500, 600, 1000, 1500, 2000,  3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11500, 13000, 14500, 16000, 17500, 19000, 20500, 22000, 23500, 25000, 27500, 30000, 32500, 35000, 37500, 40000, 42500, 45000, 47500, 50000, 55000, 60000, 65000, 70000, 75000, 80000, 85000, 90000, 95000, 100000, 110000, 120000, 130000, 140000, 150000, 160000, 170000, 180000, 190000, 200000, 210000, 220000, 230000, 240000, 250000, 260000, 270000, 280000, 290000, 300000, 310000, 320000, 330000, 340000, 350000, 360000, 370000, 380000, 390000, 400000, 410000, 420000, 430000, 440000, 450000, 460000, 470000, 480000, 490000, 500000, 510000, 520000, 530000, 540000, 550000, 560000, 570000, 580000, 590000, 600000}

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
	Species *PokemonSpecies
	Height  int
	Weight  int
	Gender  string

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

func (s *PokemonSheet) Write(){
  general.SetJsonData(sheeters.SHEETSPATH + strconv.Itoa(s.Id) + "_" + strconv.Itoa(POKEMON_SHEETID) + ".json", s)
}

func (s *PokemonSheet) LvlUp(x int){
  s.Lvl += x
  s.Status.Distributable[1] += x

  s.ElemBonus = s.Lvl/5

  general.SetJsonData("data/sheets/"+strconv.Itoa(s.Id)+"_1.json", s)
}

func (s *PokemonSheet) CalcEvasions(){
  s.Evasion[0] = general.Capped(s.Status.Total["DEF"] / 5, 0, 6)
  s.Evasion[1] = general.Capped(s.Status.Total["SPDEF"] / 5, 0, 6)
  s.Evasion[2] = general.Capped(s.Status.Total["SPD"] / 5, 0, 6)
}

//Calculates the pokemon's total stats (accounting stages) and elemental bonus
func (s *PokemonSheet) CalcStats(){
  for stat, val := range s.Status.Total{
    if s.Status.Stages[stat] < 0{
      s.Status.Total[stat] = val + ((val/10) * (s.Status.Stages[stat]))
    }else{
      s.Status.Total[stat] = val + ((val/4) * (s.Status.Stages[stat]))
    }
  }

  s.ElemBonus = s.Lvl/5
}

func (s *PokemonSheet) CalcHp(){
  s.Hp[1] = (s.Lvl + s.Status.Total["HP"]) * 4
  s.Hp[0] = s.Hp[1]
}

func (s *PokemonSheet) AllocateStats(vector map[string]int){
  var allocSum int

  for key, val := range vector{
    s.Status.LvlUp[key] += val
    allocSum += val
  }

  s.Status.Distributable[0] += allocSum

  s.CalcStats()
  s.CalcEvasions()
  s.CalcHp()

  general.SetJsonData("data/sheets/"+strconv.Itoa(s.Id)+"_1.json", s)
}
func (s *PokemonSheet) Update(nickname string, hp, atkStage, defStage, spatkStage, spdefStage, spdStage int, notes string){
  s.Nick = nickname

  s.Hp[0] = hp
  s.Status.Stages["ATK"] = atkStage
  s.Status.Stages["DEF"] = defStage
  s.Status.Stages["SPATK"] = spatkStage
  s.Status.Stages["SPDEF"] = spdefStage
  s.Status.Stages["SPD"] = spdStage

  s.Notes = notes

  s.CalcStats()

  s.Write()
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

	err = template.Execute(buf, s)
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
  Expertises  []Expertise

	TotalSeenPokemon   int
	SeenPokemon        [807]bool
	TotalCaughtPokemon int
	CaughtPokemon      [807]bool

	Inventory []*Item

	PokemonList []*PokemonSheet

	Prizes string

	Notes string
}
func (s *TrainerSheet) Write(){
  general.SetJsonData(sheeters.SHEETSPATH + strconv.Itoa(s.Id) + "_" + strconv.Itoa(TRAINER_SHEETID) + ".json", s)
}

func (s *TrainerSheet) AddExpertise(ex *Expertise){
  s.Expertises = append(s.Expertises, *ex)
}

func (s *TrainerSheet) AddClass(class *TrainerClass){
  i := 0
  for i = 0; s.Classes[i] != nil; i++{}
  s.Classes[i] = class

  s.TalentSlots--

  s.Talents = append(s.Talents, class.BasicTalents[0])
  s.Talents = append(s.Talents, class.BasicTalents[1])
}

func (s *TrainerSheet) AddTalent(talent *TrainerTalent){
  s.Talents = append(s.Talents, talent)

  s.TalentSlots--
}

func (s *TrainerSheet) GeneralUpdate(f url.Values) error{
  return nil
}

func (s *TrainerSheet) Render(w http.ResponseWriter) error {
	tmpl, err := sheeters.RenderVolatile("sheet.page.html", "./ui/html")
	if err != nil {
		return err
	}

	tmpl, err = tmpl.ParseFiles("./pkg/pokemon/PTA1/html/trainerTalentBox.partial.html")
	tmpl, err = tmpl.ParseFiles("./pkg/pokemon/PTA1/html/trainerSheet.partial.html")

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, s)
	if err != nil {
		return err
	}

	buf.WriteTo(w)

	return nil
}

func (s *TrainerSheet) LvlUp(x int){
  s.Lvl += x
  s.Status.Distributable[1] = TRAINERLVLTABLE["total_status"][s.Lvl + x]
  s.TalentSlots += TRAINERLVLTABLE["talents"][s.Lvl + x]

  general.SetJsonData("data/sheets/"+strconv.Itoa(s.Id)+"_0.json", s)
}

func (s *TrainerSheet) CalcEvasions(){
  s.Evasion[0] = general.Capped(s.Status.Total["DEF"] / 5, 0, 6)
  s.Evasion[1] = general.Capped(s.Status.Total["SPDEF"] / 5, 0, 6)
  s.Evasion[2] = general.Capped(s.Status.Total["SPD"] / 5, 0, 6)
}

//Calculates the trainer's stat modifiers, total stats (accounting stages) and movements
func (s *TrainerSheet) CalcStats(){
  for stat, val := range s.Status.Status{
    if val < 10{
      s.Status.Modifiers[stat] = val - 10
    }else{
      s.Status.Modifiers[stat] = (val-10)/2
    }
    
    if s.Status.Stages[stat] < 0{
      s.Status.Total[stat] = val + ((val/10) * (s.Status.Stages[stat]))
    }else{
      s.Status.Total[stat] = val + ((val/4) * (s.Status.Stages[stat]))
    }
  }

  s.Movement["land"] = general.Capped(3 + general.Capped(s.Status.Modifiers["ATK"]/2, s.Status.Modifiers["SPD"]/2, 1000), 5, 1000)
  s.Movement["swim"] = general.Capped(2 + s.Status.Modifiers["DEF"], 4, 1000)
  if s.Status.Modifiers["ATK"] >= 3 || s.Status.Modifiers["DEF"] >=3 {
    s.Movement["underwater"] = 4
  }else{
    s.Movement["underwater"] = 3
  }
}

func (s *TrainerSheet) CalcHp(){
  s.Hp[1] = (s.Lvl + s.Status.Status["HP"]) * 4
  s.Hp[0] = s.Hp[1]
}

func (s *TrainerSheet) AllocateStats(vector map[string]int){
  var allocSum int

  for key, val := range vector{
    s.Status.Status[key] += val
    allocSum += val
  }

  s.Status.Distributable[0] += allocSum

  s.CalcStats()
  s.CalcEvasions()
  s.CalcHp()

  general.SetJsonData("data/sheets/"+strconv.Itoa(s.Id)+"_0.json", s)
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

	Frequency string
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
