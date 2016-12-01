package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"math/rand"
	"time"
	"fmt"
	"strings"
)

const (
	FirstArg = iota
	OFlag
	OutputFile
)

func PrintHelp() {
	fmt.Println("Brick Dataset Generator ©2016 ElectricCoffee.")
	fmt.Println("Usage:\n\n\tbrick-generator generator [-o output-file]\n")
	fmt.Println("generator is a JSON file " +
		"containing the required generator parameters.\n")
	fmt.Println("-o output-file tells the program " +
		"to output to the given file instead of stdout.")
}

func main() {
	rand.Seed(time.Now().Unix())
	arguments := os.Args[1:] // arguments without program name
	argLen := len(arguments)

	if argLen < 1 {
		fmt.Println("Please supply a generator file.")
		return
	}

	if argLen == 1 {
		arg := arguments[FirstArg]
		// if arg 1 is -h or --help, write the list of commands and an introduction
		if strings.Compare(arg, "--help") == 0 ||
			strings.Compare(arg, "-h") == 0 {
			PrintHelp()
		} else {	// else write the file to stdout
			dataSet := FileToDataSet(arg)
			fmt.Println(SWriteDataSet(dataSet))
		}
		
		return
	} else if argLen == 3 && strings.Compare(arguments[OFlag], "-o") == 0 {
		inputFileName  := arguments[FirstArg]
		outputFileName := arguments[OutputFile]

		dataSet := FileToDataSet(inputFileName)
		err := FWriteDataSet(outputFileName, dataSet)

		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Invalid arguments,", arguments)
	}
}
