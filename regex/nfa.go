package regex

// NFA data structures

// State state data structure of NFA Atomaton
type State struct {
    symbol rune
    edge1 *State
    edge2 *State
}

// Nfa Auxiliary data structure
type Nfa struct {
    initial *State
    accept *State 
}

// Specials Global map that contains the special symbols
 var Specials = map[rune]int{'*': 10, '?': 9, '.': 8, '|': 7}