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
	transitions map[string]State
}

// Global variables
var vocab string
var n int
var states map[string]State

func main() {
	fmt.Println("Enter Vocabulary")
	fmt.Scanf("%s\n", &vocab)
	fmt.Print("Enter number of states\n")
	fmt.Scanf("%d\n", &n)
	states = make(map[string]State)
	states = initStates()
	inputTransitions()
	fmt.Println(states)
}

func initStates() map[string]State {
	states := make(map[string]State)
	for i := 0; i < n; i++ {
		states["q"+strconv.Itoa(i+1)] = State{
			name:        "q" + strconv.Itoa(i+1),
			isFinal:     false,
			isInitial:   false,
			transitions: make(map[string]State),
		}
	}
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
			// fmt.Scanf("%s %s\n", &inpString, &destinationState)
			fmt.Scanf("%s\n", &inpString)
			fmt.Scanf("%s\n", &destinationState)
			//TODO: Check if input string and destination state is valid before assigning
			states[getStateNameFromID(i)].transitions[inpString] = states[destinationState]
			fmt.Println(states)
		}
	}
}

func getStateNameFromID(i int) string {
	return "q" + strconv.Itoa(i+1)
}
