package main

import (
	"fmt"
)

func minimize() {
	mat = make([][]bool, n)
	for i := range mat {
		mat[i] = make([]bool, n)
	}
	// Rows start from q1 and end at qn-1, columns start from q0 and end at qn-2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				mat[i][j] = true // Never consider the upper triangle of the matrix
			}
		}
	}
	// Final and non-final states will never be equivalent
	for i := 1; i <= n-1; i++ {
		for j := 0; j <= n-2; j++ {
			if states[getStateNameFromID(i)].isFinal && !states[getStateNameFromID(j)].isFinal {
				if i > j {
					mat[i][j] = true
				} else {
					mat[j][i] = true
				}
			}
		}
	}
	// Iterate until no pair of states can be crossed out
	var flag bool
	for {
		flag = false
		for i := 1; i <= n-1; i++ {
			for j := 0; j <= n-2; j++ {
				if !mat[i][j] {
					state1 := states[getStateNameFromID(i)]
					state2 := states[getStateNameFromID(j)]
					for k := range inputSymbols {
						id1, _ := getIDfromStateName(state1.transitions[k].name)
						id2, _ := getIDfromStateName(state2.transitions[k].name)
						if id1 == id2 {
							continue
						}
						if ((id1 > id2) && (mat[id1][id2])) || (id1 < id2 && mat[id2][id1]) {
							mat[i][j] = true
							flag = true
							break
						}
					}
				}
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(mat)
	for i := 1; i <= n-1; i++ {
		for j := 0; j <= n-2; j++ {
			if !mat[i][j] {
				fmt.Printf("%s and %s are non-distinguishable\n", getStateNameFromID(i), getStateNameFromID(j))
				states[getStateNameFromID(i)+","+getStateNameFromID(j)] = &State{
					name:        getStateNameFromID(i) + "," + getStateNameFromID(j),
					isInitial:   states[getStateNameFromID(i)].isInitial,
					isFinal:     states[getStateNameFromID(i)].isFinal,
					transitions: states[getStateNameFromID(i)].transitions,
				}
				delete(states, getStateNameFromID(i))
				delete(states, getStateNameFromID(j))
			}
		}
	}
	fmt.Println("\nNew transition table:")
	printTransitionTable()
}
