package main

import "fmt"

func initialProcessing() {
	mat = make([][]bool, n)
	for i := range mat {
		mat[i] = make([]bool, n)
	}
	// Final and non-final states will never be equivalent
	// Also the same state appearing in both row and column should not cause it to be removed
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i == j) || (states[getStateNameFromID(i)].isFinal && !states[getStateNameFromID(j)].isFinal) {
				mat[i][j] = true
				mat[j][i] = true
			}
		}
	}
	fmt.Println(mat)
}
