package main

import (
	"fmt"
	"strconv"
)

// State represents a single state in the DFA
type State struct {
	name        string
	isInitial   bool
	isFinal     bool
	transitions map[string]*State
}

// Global variables
var inputSymbols map[string]bool
var n int
var states map[string]*State
var initialState *State
var mat [][]bool

func main() {
	fmt.Println("Enter input symbols as a string (Each character is treated as an individual input symbol)")
	var inpSymb string
	inputSymbols = make(map[string]bool)
	fmt.Scanf("%s\n", &inpSymb)
	for _, ch := range inpSymb {
		inputSymbols[string(ch)] = true
	}
	fmt.Print("Enter number of states\n")
	fmt.Scanf("%d\n", &n)
	states = make(map[string]*State)
	states = initStates()
	inputTransitions()
	getInitialState()
	getFinalStates()
	// fmt.Println(states)
	minimize()
}

func getIDfromStateName(name string) (int, error) {
	return strconv.Atoi(name[1:])
}

func getStateNameFromID(i int) string {
	return "q" + strconv.Itoa(i)
}
