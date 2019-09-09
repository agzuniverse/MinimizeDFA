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
var inputSymbols map[rune]bool
var n int
var states map[string]*State

func main() {
	fmt.Println("Enter input symbols as a string (Each character is treated as an individual input symbol)")
	var inpSymb string
	inputSymbols := make(map[rune]bool)
	fmt.Scanf("%s\n", &inpSymb)
	for _, ch := range inpSymb {
		inputSymbols[ch] = true
	}
	fmt.Print("Enter number of states\n")
	fmt.Scanf("%d\n", &n)
	states = make(map[string]*State)
	states = initStates()
	inputTransitions()
}

func initStates() map[string]*State {
	states := make(map[string]*State)
	fmt.Print("The states are: ")
	for i := 0; i < n; i++ {
		states[getStateNameFromID(i)] = &State{
			name:        getStateNameFromID(i),
			isFinal:     false,
			isInitial:   false,
			transitions: make(map[string]*State),
		}
		fmt.Print(getStateNameFromID(i) + ", ")
	}
	fmt.Println()
	return states
}

func inputTransitions() {
	for i := 0; i < n; i++ {
		fmt.Printf("Enter number of state transitions from %s\n", getStateNameFromID(i))
		var k int
		fmt.Scanf("%d\n", &k)
		fmt.Printf("Enter %d transitions as <input string> <destination state>\n", k)
		var inpString string
		var destinationState string
		for j := 0; j < k; j++ {
			fmt.Scanf("%s %s\n", &inpString, &destinationState)
			//TODO: Check if input string is valid
			_, stateExists := states[destinationState]
			if stateExists {
				states[getStateNameFromID(i)].transitions[inpString] = states[destinationState]
			} else {
				fmt.Println("Invalid state")
				j--
			}
		}
	}
}

func getStateNameFromID(i int) string {
	return "q" + strconv.Itoa(i+1)
}
