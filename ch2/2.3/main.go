package main


import (
	"fmt"

	"./popcount"
)

func main() {
	const num = 4354

	fmt.Println(popcount.PopCount(num))
	fmt.Println(popcount.WithLoop(num))
}