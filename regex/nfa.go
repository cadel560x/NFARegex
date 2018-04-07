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