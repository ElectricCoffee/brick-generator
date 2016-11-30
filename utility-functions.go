// utility-functions.go contains various functions that deal with rudimentary
// conversion and convenience functions
package main

import (
	"fmt"
	"strconv"
	"path/filepath"
	"strings"
	"errors"
)

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

func MakeOutputFileName(inputFileName string) (string, error) {
	var fileExt, noSuffix string
	fileExt = filepath.Ext(inputFileName)

	if strings.Compare(fileExt, ".json") != 0 {
		return "", errors.New("Please make sure the supplied file is a .json file")
	}

	noSuffix = strings.TrimSuffix(inputFileName, fileExt)
	
	return fmt.Sprint(noSuffix, ".brk"), nil
}
