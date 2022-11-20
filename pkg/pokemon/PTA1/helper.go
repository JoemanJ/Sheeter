package PTA1

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Saves the unmarshalled content of json file "path" to the variable pointed by "m"
func getJsonData(path string, m interface{}) error {
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
func setJsonData(path string, m interface{}) error {
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
