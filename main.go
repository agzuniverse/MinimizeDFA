package main

import (
	"fmt"
)

// State represents a single state in the DFA
type State struct {
	name        string
	isInitial   bool
	isFinal     bool
	transitions map[string]*State
}

// Global variables
var inputSymbols map[rune]bool
var n int
var states map[string]*State

func main() {
	fmt.Println("Enter input symbols as a string (Each character is treated as an individual input symbol)")
	var inpSymb string
	inputSymbols = make(map[rune]bool)
	fmt.Scanf("%s\n", &inpSymb)
	for _, ch := range inpSymb {
		inputSymbols[ch] = true
	}
	fmt.Print("Enter number of states\n")
	fmt.Scanf("%d\n", &n)
	states = make(map[string]*State)
	states = initStates()
	inputTransitions()
	getInitialState()
	getFinalStates()
}
