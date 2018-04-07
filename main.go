package main

import (
	"fmt"

	"./regex"
)

func main() {
	fmt.Println(regex.Pomatch("ab.c*|", "ab"))
	fmt.Println(regex.Pomatch("ab.c*|", "a"))
	fmt.Println(regex.Pomatch("ab.c*|", "b"))
	fmt.Println(regex.Pomatch("ab.c*|", ""))
	fmt.Println(regex.Pomatch("ab.c*|", "c"))
	fmt.Println(regex.Pomatch("ab.c*|", "ccc"))
}