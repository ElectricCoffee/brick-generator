// brick.go contains functions related to the Brick struct
package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

type Brick struct {
	Colour uint
	Size uint
}

func NewBrick(colour, length uint) Brick {
	return Brick{colour, length}
}

func (b Brick) String() string {
	return fmt.Sprintf("COL:%d LEN:%d\n", b.Colour, b.Size)
}
// SWriteDataSet takes a dataset of type []Brick and
// constructs a string in the format required by the output file
// The name follows C's naming convention with an s prefix indicating a string output
func SWriteDataSet(input []Brick) string {
	strArr := make([]string, len(input))

	for i, brick := range input {
		strArr[i] = brick.String()
	}

	return strings.Join(strArr, "")
}

// WriteDataSet takes a dataset of type []Brick and
// writes it to standard output in the format required by the output file
// The name follows C's naming convention with no prefix indicating stdout
func WriteDataSet(input []Brick) {
	fmt.Println(SWriteDataSet(input))
}

// FWriteDataSet takes a dataset of type []Brick and
// writes a file in the format required by the output file
// The name follows C's naming convention with an f prefix indicating a file output
func FWriteDataSet(filename string, input []Brick) error {
	byteArr := []byte(SWriteDataSet(input))
	// 0644 is the unix permission that sets -rw-r--r--
	return ioutil.WriteFile(filename, byteArr , 0644)
}
