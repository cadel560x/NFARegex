package regex


// Shunting algorithm based on video
// https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e


// import (
// 	"fmt"
// )
    
// Intopost converts from infix notation to postfix notation
func Intopost(infix string) string {

	pofix, s := []rune{}, []rune{}

	for _, r := range infix{
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
		for s[len(s)-1] != '(' {
			pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
		}

		s = s[:len(s)-1]
		case Specials[r] > 0:
			for len(s) > 0 && Specials[r] <= Specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)
		default:
			pofix = append(pofix, r)

		}
	}

	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)

}


// func main()  {
// 	fmt.Println("Infix:  ", "a.b.c*")
// 	fmt.Println("Postfix:  ", intopost("a.b.c*"))

// 	fmt.Println("Infix:  ", "(a.(b|d))*")
// 	fmt.Println("Postfix:  ", intopost("(a.(b|d))*"))

// 	fmt.Println("Infix:  ", "a.(b|d).c*")
// 	fmt.Println("Postfix:  ", intopost("a.(b|d).c*"))

// 	fmt.Println("Infix:  ", "a.(b.b)+.c")
// 	fmt.Println("Postfix:  ", intopost("a.(b.b)+.c"))
// }