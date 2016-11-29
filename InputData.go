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
	var parsedData InputData
	inputFile, readErr := ioutil.ReadFile(fileName)

	if readErr != nil {
		return parsedData, readErr
	}

	// Extracts JSON data into the variable parsedData
	jsErr := json.Unmarshal(inputFile, &parsedData)

	return parsedData, jsErr
}
