package storage

import (
	"encoding/json"
	"os"
)

func SaveJSON(filename string, data any) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0o644)
}

func LoadJSON[T any](filename string) (T, error) {
	var out T
	data, err := os.ReadFile(filename)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(data, &out)

	return out, err
}
