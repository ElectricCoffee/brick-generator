// data-generation.go contains the functions responsible for generating the output data.
// This includes generating the randomised data when it's needed
package main

import (
	"math"
	"math/rand"
)


// Constant values for the colours black and white
const (
	Black = 0x000000
	White = 0xFFFFFF
	NumberOfBricks = 10
)

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

// GenerateSizes generates a sequence of pseudorandom entries from the
// BrickSizes array
func GenerateSizes(number uint) []uint {
	var result []uint
	for i := uint(0); i < number; i++ {
		l := uint(rand.Intn(10) + 1)
		result = append(result, l)
	}

	return result
}

// ScaleBrickSlice takes a brick array pointer and a max length,
// and scales the underlying array to have its length match the max length.
// It does so by either removing n excess items,
// or duplicating the first n items to fill out the slice.
func ScaleBrickSlice(brksPtr *[]Brick, maxLen uint) {
	brickLen := uint(len(*brksPtr))

	if brickLen > maxLen {
		// if the size of the brksPtr slice is longer than maxlen
		// then scale it down
		*brksPtr = append([]Brick(nil), (*brksPtr)[:maxLen]...)
	} else if brickLen < maxLen {
		lengthDifference := maxLen - brickLen
		if brickLen < maxLen / uint(2) {
			// if it's shorter, check if it's shorter than half the length
			// append the array to itself, and call the function again
			*brksPtr = append(*brksPtr, *brksPtr...)
			ScaleBrickSlice(brksPtr, maxLen)
		} else {
			// if it isn't, append the difference from itself and return
			appendValue := make([]Brick, lengthDifference)
			copy(appendValue, (*brksPtr)[:lengthDifference])
			*brksPtr = append(*brksPtr, appendValue...)
		}
	} // else do nothing
}

// FillInMissing generates additional colours or brick sizes if any of them are amiss.
func FillInMissing(colours, sizes *[]uint, maxNum uint) {
	noColours := len(*colours) == 0
	noSizes := len(*sizes) == 0

	calcSize := func (slice []uint) uint {
		return maxNum / uint(len(slice))
	}

	if noColours && !noSizes {
		colCard := calcSize(*sizes)
		*colours  = GenerateColours(colCard)
	} else if !noColours && noSizes {
		sizCard := calcSize(*colours)
		*sizes  = GenerateSizes(sizCard)
	} else if noColours && noSizes {
		// if neither colours nor sizes are present,
		// generate the root maxNum of each
		cardinality := uint(math.Sqrt(float64(maxNum)))
		*colours  = GenerateColours(cardinality)
		*sizes  = GenerateSizes(cardinality)
	}
}

// GenerateDataSet creates a sequence of bricks based on the input slices.
// If either colours or sizes are empty, random values will be provided.
func GenerateDataSet(colours, sizes []uint, maxNum uint) []Brick {
	var bricks []Brick

	FillInMissing(&colours, &sizes, maxNum)
	
	// generate cartesian product of colours and sizes
	for _, colour := range colours {
		for _, size := range sizes {
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
