package regex

// Shunting algorithm based on video
// https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b

import(
    "fmt"
)

// PostfixRegexNFA From postfix regex to NFA
func PostfixRegexNFA(postfix string) *Nfa {
    nfastack := []*Nfa{}

    for _, r := range postfix {

        switch r {
        case '.':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            frag1.accept.edge1 = frag2.initial

            nfastack = append(nfastack, &Nfa{initial: frag1.initial, accept: frag2.accept})
            
        case '|':
            frag2 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]
            
            initial := State{edge1: frag1.initial, edge2: frag2.initial}
            accept := State{}
            frag1.accept.edge1 = &accept
            frag2.accept.edge1 = &accept            

            nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})
        
        case '*':
            frag := nfastack[len(nfastack)-1]
            nfastack := nfastack[:len(nfastack)-1]

            accept := State{}
            initial := State{edge1: frag.initial, edge2: &accept}
            frag.accept.edge1 = frag.initial
            frag.accept.edge2 = &accept

            nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})

        default: 
            accept := State{}
            initial := State{symbol: r, edge1: &accept}

            nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})
        
        }
    }
    
    if len(nfastack) != 1 {
		fmt.Println("Sorry more than 1 nfa found", len(nfastack), nfastack)
    }

    return nfastack[0]

}

// func main(){
//     nfa := postfixRegexNFA("ab.c*|")
//     fmt.Println(nfa)
// }