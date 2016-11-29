package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"fmt"			// imports Printf and cousins
	"strconv"		// handles string conversion
)



func main() {
	var inputFileName, outputFileName string
	var colours []uint

	arguments := os.Args[1:] // arguments without program name

	if len(arguments) < 1 {
		panic("Please supply a generator file")
	}

	inputFileName = arguments[0] // The JSON file we need to read

	parsedData, jsErr := DataFromFile(inputFileName)

	if jsErr != nil {
		panic(readErr)
	}

	// Populate the uint array colours with the converted values
	for _, value := range parsedData.Colours {
		colour, convErr := strconv.ParseUint(value, 0, 32)
		if convErr != nil {
			panic(convErr)
		}
		
		// ParseUint returns a uint64, we need a 32, hence casting
		colours = append(colours, uint(colour))
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
