package regex

// Shunting algorithm based on video
// https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b

// import(
//     "fmt"
// )

type state struct {
    symbol rune
    edge1 *state
    edge2 *state
}

type nfa struct {
    initial *state
    accept *state 
}

// PostfixRegexNFA From postfix regex to NFA
func PostfixRegexNFA(postfix string) *nfa{
    nfastack := []*nfa{}

    for _, r := range postfix {

        switch r {
        case '.':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            frag1.accept.edge1 = frag2.initial

            nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
            
        case '|':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            initial := state{edge1: frag1.initial, edge2: frag2.initial}
            accept := state{}
            frag1.accept.edge1 = &accept
            frag2.accept.edge2 = &accept
            

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        
        case '*':
            frag := nfastack[len(nfastack)-1]
            nfastack := nfastack[:len(nfastack)-1]

            accept := state{}
            initial := state{edge1: frag.initial, edge2: &accept}
            frag.accept.edge1 = frag.initial
            frag.accept.edge2 = &accept

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

        default: 
            accept := state{}
            initial := state{symbol: r, edge1: &accept}

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        
        }//switch
    }// for
    return nfastack[0]
}

// func main(){
//     nfa := postfixRegexNFA("ab.c*|")
//     fmt.Println(nfa)
// }