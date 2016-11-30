package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"fmt"			// imports Printf and cousins
	"math/rand"
	"time"
	"path/filepath"
	"strings"
	"errors"
)

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
