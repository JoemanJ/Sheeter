package sheeters

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

const RDPATH string = "./data/RD.json"
const SHEETSPATH string = "./data/sheets/"

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

	fileName := filepath.Join(SHEETSPATH, strconv.Itoa(id)+"_sheet.json")

	err := GetJsonData(fileName, &sheet)
	if err != nil {
		return nil, err
	}

	return &sheet, nil
}

// Saves the unmarshalled content of json file "path" to the variable pointed by "m"
func GetJsonData(path string, m interface{}) error {
	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		json.Unmarshal([]byte("{}"), m)
		SetJsonData(path, m)
		return err
	}

	content, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, m)
	if err != nil {
		return err
	}

	return nil
}

// Saves the content of "m" on a json file on path "path"
func SetJsonData(path string, m interface{}) error {
	jsonText, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		s := fmt.Sprintf("Error marshalling file %s:\n%s", path, err.Error())
		return errors.New(s)
	}

	err = os.WriteFile(path, jsonText, 0755)
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
	var file []byte

	_, err := os.Stat(RDPATH)

	if errors.Is(err, os.ErrNotExist) { //FILE DOESNT EXIST: DATA = EMPTY MAP
		data = map[string]string{}
	} else {
		file, err = os.ReadFile(RDPATH)
		if err != nil { //RANDOM ERROR: RETURN IT
			return err
		}
		//FILE EXISTS: DATA = FILE
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

	err = os.WriteFile(RDPATH, jsonTxt, 0755)
	if err != nil {
		return err
	}

	return nil
}
