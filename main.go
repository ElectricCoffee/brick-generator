package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"fmt"			// imports Printf and cousins
	"strconv"		// handles string conversion
	"math/rand"
	"time"
	"path/filepath"
	"strings"
	"errors"
)

// Constant values for the colours black and white
const (
	Black = 0x000000
	White = 0xFFFFFF
)

// BrickLengths uses the index as the length of a brick in studs,
// the numbers are the average lengths of the bricks as measured by the machine
var BrickLengths = [9]uint{0, 0, 3524, 5254, 6904, 8198, 10352, 12982, 14204}

// crash first prints the error, then causes the program to panic.
// panic doesn't print the error message (apparently),
// so a println is used in conjunction with panic to
// achieve the desired functionality
func crash(err error) {
	fmt.Println(err)
	panic(err)
}

// ConvertToIntegers takes a slice of numeric strings
// and converts it to a slice of integers.
// It returns an error if a conversion wasn't successful
func ConvertToIntegers(numStrs []string, colourInts *[]uint) error {
	// Populate the uint array colours with the converted values
	for _, str := range numStrs {
		colour, convErr := strconv.ParseUint(str, 0, 32)
		
		if convErr != nil {
			return convErr
		}
		
		// ParseUint returns a uint64, we need a 32, hence casting
		*colourInts = append(*colourInts, uint(colour))
	}
	return nil
}

func main() {
	rand.Seed(time.Now().Unix())
	var inputFileName, outputFileName string
	var colours, sizes []uint
	
	arguments := os.Args[1:] // arguments without program name

	if len(arguments) < 1 {
		crash(errors.New("Please supply a generator file"))
	}

	inputFileName = arguments[0] // The JSON file we need to read

	fileExt := filepath.Ext(inputFileName)

	if strings.Compare(fileExt, ".json") != 0 {
		crash(errors.New("Please make sure the supplied file is a .json file"))
	}

	outputFileName = fmt.Sprint(strings.TrimSuffix(inputFileName, fileExt), ".brk")

	parsedData, jsErr := DataFromFile(inputFileName)

	if jsErr != nil {
		crash(jsErr)
	}

	convErr := ConvertToIntegers(parsedData.Colours, &colours)

	for _, e := range parsedData.Sizes {
		sizes = append(sizes, BrickLengths[e])
	}

	if convErr != nil {
		crash(convErr)
	}

	bricks := GenerateDataSet(colours, sizes, parsedData.Amount)

	FWriteDataSet(outputFileName, bricks)
}
