package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"fmt"			// imports Printf and cousins
	"strconv"		// handles string conversion
	"math/rand"
	"time"
)

const (
	Black = 0x000000
	White = 0xFFFFFF
)

func ConvertToIntegers(colourHexes []string, colourInts *[]uint) error {
	// Populate the uint array colours with the converted values
	for _, hex := range colourHexes {
		colour, convErr := strconv.ParseUint(hex, 0, 32)
		
		if convErr != nil {
			return convErr
		}
		
		// ParseUint returns a uint64, we need a 32, hence casting
		*colourInts = append(*colourInts, uint(colour))
	}
	return nil
}

func GenerateColours(number uint) []uint {
	var result []uint
	for i := 0; i < number; i++ {
		// generate between 0 and FFFFFF
		colour := uint(rand.Intn(White))
		result = append(result, colour)
	}

	return result
}

func GenerateLengths(number uint) [] uint {
	var result []uint
	for i := 0; i < number; i++ {
		length := uint(rand.Intn(200))
		result = append(result, length)
	}

	return result
}

func main() {
	rand.Seed(time.Now().Unix())
	var inputFileName, outputFileName string
	var colours []uint

	arguments := os.Args[1:] // arguments without program name

	if len(arguments) < 1 {
		panic("Please supply a generator file")
	}

	inputFileName = arguments[0] // The JSON file we need to read

	parsedData, jsErr := DataFromFile(inputFileName)

	if jsErr != nil {
		panic(jsErr)
	}

	convErr := ConvertToIntegers(parsedData.Colours, &colours)

	
	
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
