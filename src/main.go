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
var states = make(map[string]State)

func main() {
	fmt.Println("Enter Vocabulary")
	fmt.Scanf("%s\n", &vocab)
	fmt.Print("Enter number of states\n")
	fmt.Scanf("%d\n", &n)
	states = initStates()
}

func initStates() (states map[string]State) {
	for i := 0; i < n; i++ {
		states["q"+strconv.Itoa(i+1)] = State{
			name:        "q" + strconv.Itoa(i+1),
			isFinal:     false,
			isInitial:   false,
			transitions: make(map[string]State),
		}
	}
	return
}

func getStateNameFromID(i int) string {
	return "q" + strconv.Itoa(i+1)
}
