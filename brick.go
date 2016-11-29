package main

import "fmt"

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

