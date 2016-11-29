package main

import (
	"os"			// imports command-line args
	"fmt"			// imports Printf and cousins
	"encoding/json"		// imports JSON utilities
	"io/ioutil"		// imports file utilities
)

type InputData struct {
	Colours []string `json:"colors"`
	Sizes   []int    `json:"sizes"`
	Amount  int      `json:"amount"`
}

func main() {
	var inputFileName, outputFileName string
	var parsedData InputData

	arguments := os.Args[1:] // arguments without program name

	if len(arguments) < 1 {
		panic("Please supply a generator file")
	}

	inputFileName = arguments[0] // The JSON file we need to read

	inputFile, readErr := ioutil.ReadFile(inputFileName)

	if readErr != nil {
		panic(readErr)
	}

	jsErr := json.Unmarshal(inputFile, &parsedData)

	if jsErr != nil {
		panic(readErr)
	}
	
	var in, out InputData
	var b []byte
	
	in = InputData {
		Colors: []string{"0xFF0000", "0x00FF00", "0x0000FF"},
		Sizes: []int{50, 100, 150},
		Amount: 500,
	}

	b, err = json.Marshal(in)
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	
	err = json.Unmarshal(b, &out)

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(out)
}
