package main

import (
	"errors"
	"fmt"
	"sort"
)

//TODO: non-integer where an integer input is expected should not crash the program
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

func validateInputString(in string) (bool, error) {
	if len(in) != 1 {
		return false, errors.New("Input string must be a single character")
	}
	_, exists := inputSymbols[in]
	if !exists {
		return false, errors.New("Invalid input string")
	}
	return true, nil
}

func inputTransitions() {
	keys := make([]string, len(inputSymbols))
	q := 0
	for key := range inputSymbols {
		keys[q] = key
		q++
	}
	sort.Strings(keys)
	for i := 0; i < n; i++ {
		var destinationState string
		for _, k := range keys {
			fmt.Printf("Enter transition from %s when given input %s\n", getStateNameFromID(i), k)
			for {
				fmt.Scanf("%s\n", &destinationState)
				_, stateExists := states[destinationState]
				if stateExists {
					states[getStateNameFromID(i)].transitions[k] = states[destinationState]
					break
				} else {
					fmt.Println("Invalid state")
				}
			}
		}
	}
}

func getInitialState() {
	fmt.Println("Enter inital state")
	var inpState string
	for {
		fmt.Scanf("%s\n", &inpState)
		state, exists := states[inpState]
		if !exists {
			fmt.Println("Invalid state")
			continue
		}
		state.isInitial = true
		initialState = state
		break
	}
}

func getFinalStates() {
	var n int
	var finalState string
	fmt.Println("Enter number of final states")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("Enter %d final states\n", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &finalState)
		state, exists := states[finalState]
		if !exists {
			fmt.Println("Invalid state")
			i--
			continue
		}
		state.isFinal = true
	}
}
