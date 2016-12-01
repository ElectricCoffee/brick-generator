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

func SWriteDataSet(input []Brick) string {
	strArr := make([]string, len(input))

	for i, brick := range input {
		strArr[i] = brick.String()
	}

	return strings.Join(strArr, "")
}

func FWriteDataSet(filename string, input []Brick) error {
	byteArr := []byte(SWriteDataSet(input))
	// 0644 is the unix permission that sets -rw-r--r--
	return ioutil.WriteFile(filename, byteArr , 0644)
}
