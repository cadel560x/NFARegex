##  Graph Theory Project
#### This is program  in Go programming language that builds a non-deterministic finite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “∗ ” have the special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1∗ means any number of 1’s. 

### Instruction 
- In command line or terminal type: git clone https://github.com/cadel560x/NFARegex.git
- Go to directory: NFARegex
- Execute: go build main.go
- Execute: main.exe or go run main.go

### What is Non-deterministic Finite Automaton
In NDFA, for a particular input symbol, the machine can move to any combination of the states in the machine. In other words, the exact state to which the machine moves cannot be determined. Hence, it is called Non-deterministic Automaton. As it has finite number of states, the machine is called Non-deterministic Finite Machine or Non-deterministic Finite Automaton.

#### Formal Definition of an NDFA

#### An NDFA can be represented by a 5-tuple (Q, ∑, δ, q0, F) where −

<b>Q</b> is a finite set of states.

<b>∑</b> is a finite set of symbols called the alphabets.

<b>δ</b> is the transition function where δ: Q × ∑ → 2Q

(Here the power set of Q (2Q) has been taken because in case of NDFA, from a state, transition can occur to any combination of Q states)

<b>q0</b> is the initial state from where any input is processed (q0 ∈ Q).

<b>F</b> is a set of final state/states of <b>Q (F ⊆ Q)</b>.


An NDFA is represented by digraphs called state diagram.

The vertices represent the states.
The arcs labelled with an input alphabet show the transitions.
The initial state is denoted by an empty single incoming arc.
The final state is indicated by double circles.

#### Example
Let a non-deterministic finite automaton be →

* Q = {a, b, c}
* ∑ = {0, 1}
* q0 = {a}
* F = {c}

### Project resource:
https://swtch.com/~rsc/regexp/regexp1.html<br />
https://web.microsoftstream.com/video/96e6f4cc-b390-4531-ba7f-84ad6ab01f47</br>
https://web.microsoftstream.com/video/d08f6a02-23ec-4fa1-a781-585f1fd8c69e</br>
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e</br>
https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c</br>
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b</br>
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d</br>