package storage

import (
	"encoding/json"
	"os"
)

// Salva uma string em um arquivo json
func SaveJSON(filename string, data any) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0o644)
}

// Func que abre um arquivo json e converte para um tipo de dado especificado
func LoadJSON[T any](filename string) (T, error) {
	var out T
	data, err := os.ReadFile(filename)
	if err != nil {
		return out, err
	}
	// Unmasrhal recebe o dado json e um ponteiro para a variavel que deseja receber o conteúdo do json
	// no caso abaixo: var out T
	// func possui apenas o error de retorno
	// pois ela não retorna o dado convertido, e sim altera uma variavel pre-definida para receber os dados do JSON
	err = json.Unmarshal(data, &out)

	return out, err
}
