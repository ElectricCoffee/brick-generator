package main

import (			// comments there for personal reminder
	"os"			// imports command-line args
	"fmt"			// imports Printf and cousins
	"strconv"		// handles string conversion
	"math"
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

// GenerateColours generates a sequence of pseudorandom colours between
// black #000000, and white #FFFFFF.
func GenerateColours(number uint) []uint {
	var result []uint
	for i := uint(0); i < number; i++ {
		// generate between 0 and FFFFFF
		colour := uint(rand.Intn(White + 1))
		result = append(result, colour)
	}

	return result
}

// GenerateLengths generates a sequence of pseudorandom entries from the
// BrickLengths array
func GenerateLengths(number uint) []uint {
	var result []uint
	for i := uint(0); i < number; i++ {
		l := uint(rand.Intn(8) + 1)
		result = append(result, BrickLengths[l])
	}

	return result
}

func ScaleBrickSlice(brksPtr *[]Brick, maxLen uint) {
	fmt.Println("Scaling bricks slice")
	brickLen := uint(len(*brksPtr))

	if brickLen > maxLen {
		fmt.Println("The length is longer")
		// if the size of the brksPtr slice is longer than maxlen
		// then scale it down
		*brksPtr = append([]Brick(nil), (*brksPtr)[:maxLen]...)
	} else if brickLen < maxLen {
		lengthDifference := maxLen - brickLen
		fmt.Println("difference is", lengthDifference)
		if brickLen < maxLen / uint(2) {
			fmt.Println("Length is shorter than half")
			// if it's shorter, check if it's shorter than half the length
			// append the array to itself, and call the function again
			*brksPtr = append(*brksPtr, *brksPtr...)
			ScaleBrickSlice(brksPtr, maxLen)
		} else {
			fmt.Println("Length is longer than half")
			// if it isn't, append the difference from itself and return
			appendValue := make([]Brick, lengthDifference)
			copy(appendValue, (*brksPtr)[:lengthDifference])
			*brksPtr = append(*brksPtr, appendValue...)
		}
	} // else do nothing
}

// GenerateDataSet creates a sequence of bricks based on the input slices.
// If either colours or lengths are empty, random values will be provided.
func GenerateDataSet(colours, lengths []uint, maxNum uint) []Brick {
	var bricks []Brick
	// var cardinality uint
	fmt.Println (
		"length of colours", len(colours),
		"length of lengths", len(lengths),
		"required bricks", maxNum,
	)
	noColours := len(colours) == 0
	noLengths := len(lengths) == 0

	calcSize := func (slice []uint) uint {
		return maxNum / uint(len(slice))
	}

	if noColours && !noLengths {
		colCard := calcSize(lengths)
		colours  = GenerateColours(colCard)
	} else if !noColours && noLengths {
		sizCard := calcSize(colours)
		lengths  = GenerateLengths(sizCard)
	} else if noColours && noLengths {
		// if neither colours nor lengths are present,
		// generate the root maxNum of each
		cardinality := uint(math.Sqrt(float64(maxNum)))
		colours  = GenerateColours(cardinality)
		lengths  = GenerateLengths(cardinality)
	}
	
	// generate cartesian product of colours and lengths
	for _, colour := range colours {
		for _, size := range lengths {
			bricks = append(bricks, NewBrick(colour, size))
		}
	}

	// Shuffle the slice of bricks so there's some randomness to it
	for i := range bricks {
		j := rand.Intn(i + 1)
		bricks[i], bricks[j] = bricks[j], bricks[i]
	}

	ScaleBrickSlice(&bricks, maxNum)
		
	return bricks
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
