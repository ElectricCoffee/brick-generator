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

func FWriteDataSet(filename string, input []Brick) error {
	strArr := make([]string, len(input))

	for i, brick := range input {
		strArr[i] = brick.String()
	}

	byteArr := []byte(strings.Join(strArr, ""))
	// 0644 is the unix permission that sets -rw-r--r--
	return ioutil.WriteFile(filename, byteArr , 0644)
}
