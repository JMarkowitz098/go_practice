package main

import (
	"fmt"

	"./unitconv"
)

const (
	numInches = 36
)

func main() {
	i := unitconv.Inch(numInches)
	f := unitconv.IToF(i)
	y := unitconv.IToY(i)

	fmt.Printf("%s = %s = %s", i, f, y)
}
