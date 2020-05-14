package main


import (
	"fmt"

	"./popcount"
)
const num = 4354

func main() {
	fmt.Println(popcount.PopCount(num))
	fmt.Println(popcount.WithLoop(num))
}