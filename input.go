package main

import (
	"errors"
	"fmt"
	"strconv"
)

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
	runes := []rune(in)
	_, exists := inputSymbols[runes[0]]
	if !exists {
		return false, errors.New("Invalid input string")
	}
	return true, nil
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
			inpStringIsValid, err := validateInputString(inpString)
			if !inpStringIsValid {
				fmt.Println(err)
				j--
				continue
			}
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
