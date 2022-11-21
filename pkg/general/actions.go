package sheeters

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const RDPATH string = "./data/RD.json"

type DiceSet struct {
	X   int
	N   int
	Mod int
}

// Returns the result of a pseudo-random "XdN + mod" dice roll
func Roll(X int, N int, mod int) int {
	total := 0
	for i := 0; i < X; i++ {
		total += rand.Intn(N) + 1
	}

	return total + mod
}

// Rolls a preset roll object (type DiceSet)
func (s *DiceSet) Roll() int {
	total := 0
	for i := 0; i < s.X; i++ {
		total += rand.Intn(s.N) + 1
	}

	return total + s.Mod
}

func CreateDiceSet(X int, N int, mod int) *DiceSet {
	return &DiceSet{X: X, N: N, Mod: mod}
}

func Capped(value, minCap, maxCap int) int {
	if value < minCap {
		return minCap
	}

	if value > maxCap {
		return maxCap
	}

	return value
}

func GetSheet(id int) (*G_sheet, error) {
	var sheet G_sheet

	err := GetJsonData("./data/sheets/"+strconv.Itoa(id)+"_sheet.json", &sheet)
	if err != nil {
		return nil, err
	}

	return &sheet, nil
}

// Saves the unmarshalled content of json file "path" to the variable pointed by "m"
func GetJsonData(path string, m interface{}) error {
	_, err := os.Stat(path)

	content, err := os.ReadFile(path)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if errors.Is(err, os.ErrNotExist) {
		json.Unmarshal([]byte("{}"), m)
		return nil
	}

	json.Unmarshal(content, m)

	return nil
}

// Saves the content of "m" on a json file on path "path"
func SetJsonData(path string, m interface{}) error {
	jsonText, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		s := fmt.Sprintf("Error marshalling file %s:\n%s", path, err.Error())
		return errors.New(s)
	}

	err = os.WriteFile(path, jsonText, 0666)
	if err != nil {
		s := fmt.Sprintf("Error writing file %s:\n%s", path, err.Error())
		return errors.New(s)
	}

	return nil
}

func GetRD(key string) (string, error) {
	file, err := os.ReadFile(RDPATH)
	if err != nil {
		return "", err
	}

	var data map[string]string

	err = json.Unmarshal(file, &data)
	if err != nil {
		return "", err
	}

	return data[key], nil
}

func SetRD(key, value string) error {
	var data map[string]string

	file, err := os.ReadFile("./data/RD.json")
	if err == os.ErrNotExist { //FILE DOESNT EXIST: DATA = EMPTY MAP
		data = map[string]string{}
	} else if err != nil { //RANDOM ERROR: RETURN IT
		return err
	} else { //FILE EXISTS: DATA = FILE
		err = json.Unmarshal(file, &data)
		if err != nil {
			return err
		}
	}

	data[key] = value

	jsonTxt, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(RDPATH, jsonTxt, 0666)
	if err != nil {
		return err
	}

	return nil
}
