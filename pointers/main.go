package main

import(
	"fmt"
)

func main () {
	x := 5
	fmt.Println(x)
	p := &x
	*p = 10
	fmt.Println(x)
	fmt.Println(*p)

}