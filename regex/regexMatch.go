package regex

// Shunting algorithm based on video
// https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

// import(
//     "fmt"
// )

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