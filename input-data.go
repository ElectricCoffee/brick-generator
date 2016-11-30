package main

import (
	"encoding/json"		// imports JSON utilities
	"io/ioutil"		// imports file utilities
)

type InputData struct {
	Colours []string `json:"colors"`
	Sizes   []uint   `json:"sizes"`
	Amount  uint     `json:"amount"`
}

func DataFromFile(fileName string) (InputData, error) {
	inputFile, readErr := ioutil.ReadFile(fileName)

	if readErr != nil {
		return InputData{}, readErr
	}

	// Extracts JSON data into the variable parsedData

	return DataFromJSON(inputFile)
}

func DataFromJSON(jsonStr []byte) (InputData, error) {
	var parsedData InputData
	jsErr := json.Unmarshal(jsonStr, &parsedData)

	return parsedData, jsErr
}
