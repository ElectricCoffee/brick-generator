package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"math/rand"
	"time"
	"errors"
)

func main() {
	rand.Seed(time.Now().Unix())
	var inputFileName, outputFileName string
	var err error
	var colours, sizes []uint
	
	arguments := os.Args[1:] // arguments without program name

	if len(arguments) < 1 {
		crash(errors.New("Please supply a generator file"))
	}

	inputFileName = arguments[0] // The JSON file we need to read

	outputFileName, err = MakeOutputFileName(inputFileName)

	if err != nil {
		crash(err)
	}

	parsedData, err := DataFromFile(inputFileName)

	if err != nil {
		crash(err)
	}

	err = ConvertToIntegers(parsedData.Colours, &colours)

	if err != nil {
		crash(err)
	}
	
	for _, e := range parsedData.Sizes {
		sizes = append(sizes, BrickLengths[e])
	}

	bricks := GenerateDataSet(colours, sizes, parsedData.Amount)

	FWriteDataSet(outputFileName, bricks)
}
