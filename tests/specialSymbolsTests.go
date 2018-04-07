package main

import (
	"fmt"
	"../regex"
)

func main() {
	fmt.Println("Pomatch test")
	fmt.Print("'ab.c*|', 'ab'")
	fmt.Println(regex.Pomatch("ab.c*|", "ab"))
	fmt.Print("'ab.c*|', 'a'")
	fmt.Println(regex.Pomatch("ab.c*|", "a"))
	fmt.Print("'ab.c*|', 'b'")
	fmt.Println(regex.Pomatch("ab.c*|", "b"))
	fmt.Print("'ab.c*|', ''")
	fmt.Println(regex.Pomatch("ab.c*|", ""))
	fmt.Print("'ab.c*|', 'c'")
	fmt.Println(regex.Pomatch("ab.c*|", "c"))
	fmt.Print("'ab.c*|', 'ccc'")
	fmt.Println(regex.Pomatch("ab.c*|", "ccc"))

	fmt.Println()

	fmt.Println("RegexNFA tests")
	fmt.Print("'a.b|c*','ab' ")
	fmt.Println(regex.RegexNFA("a.b|c*","ab"))
	fmt.Print("'a.b|c*','a' ")
	fmt.Println(regex.RegexNFA("a.b|c*","a"))
	fmt.Print("'a.b|c*','b' ")
	fmt.Println(regex.RegexNFA("a.b|c*","b"))
	fmt.Print("'a.b|c*','' ")
	fmt.Println(regex.RegexNFA("a.b|c*",""))
	fmt.Print("'a.b|c*','c' ")
	fmt.Println(regex.RegexNFA("a.b|c*","c"))
	fmt.Print("'a.b|c*','ccc' ")
	fmt.Println(regex.RegexNFA("a.b|c*","ccc"))

	fmt.Println()

	fmt.Println("'+' special character")
	fmt.Print("'c+','c' ")
	fmt.Println(regex.RegexNFA("c+","c"))

	fmt.Println()

	fmt.Println("'?' special character")
	fmt.Print("'c?','c' ")
	fmt.Println(regex.RegexNFA("c?","c"))

	fmt.Print("'c?','' ")
	fmt.Println(regex.RegexNFA("c?",""))

	fmt.Print("'c?|a.b','ab' ")
	fmt.Println(regex.RegexNFA("c?|a.b","ab"))

	fmt.Print("'a.b.c?','abc' ")
	fmt.Println(regex.RegexNFA("a.b.c?","abc"))

	fmt.Print("'c?|a.b','ab' ")
	fmt.Println(regex.RegexNFA("c?|a.b","ab"))

	fmt.Print("'c?','cc' ")
	fmt.Println(regex.RegexNFA("c?","cc"))

	fmt.Print("'c?','ab' ")
	fmt.Println(regex.RegexNFA("c?","ab"))

	fmt.Print("'a.b|c?','abc' ")
	fmt.Println(regex.RegexNFA("a.b|c?","abc"))
}