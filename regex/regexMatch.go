package regex

// Shunting algorithm based on video
// https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

import(
	"fmt"
	"time"
)

// ToInfix Adds concatenation symbol to consecutive regular characters
func ToInfix(regex string) string {
	infix := []rune{}

	for i, r := range regex {
		infix = append(infix, r)
		if (i < len(regex)-1 ) {
			if (Specials[r] == 0) && (string(r) != "(" ) {
				if (Specials[rune(regex[i+1])] == 0) && ( string(regex[i+1]) != ")" ) {
					infix = append(infix, '.')
				}
			} 
		}
	}

	return string(infix)

}

// RegexNFA wrapper that converts infix to postfix and passes it to 'pomatch'
func RegexNFA(regexString string, language string) bool {
	infixRegex := ToInfix(regexString)
	postfixRegex := Intopost(infixRegex)

	// Calcultaes the time elapsed by 'Pomatch' function
	start := time.Now()
	isMatch := Pomatch(postfixRegex, language)
	elapsed := time.Since(start)

	fmt.Printf("Pomatch took %s\n", elapsed)
	
	return isMatch
}

func addState(l []*State, s *State, a *State) []*State {
    l = append(l, s)
    if s != a && s.symbol == 0 {
        l = addState(l, s.edge1, a)
        if s.edge2 != nil {
            l = addState(l, s.edge2, a)
        }
    }
    return l
}

// Pomatch Evaluates a string against a postfix regular expression
func Pomatch(po string, s string) bool {
	isMatch := false
	ponfa := PostfixRegexNFA(po)

	current := []*State{}
	next := []*State{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {		
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*State{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			isMatch = true
			break
		}
	}
	return isMatch

}

// func main(){
//     fmt.Println(Pomatch("ab.c*|", "ab"))
// }