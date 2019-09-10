package main

import "fmt"

func minimize() {
	mat = make([][]bool, n)
	for i := range mat {
		mat[i] = make([]bool, n)
	}
	// Rows start from q1 and end at qn-1, columns start from q0 and end at qn-2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				mat[i][j] = true
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
	fmt.Println(mat)
	for i := 1; i <= n-1; i++ {
		for j := 0; j <= n-2; j++ {
			if !mat[i][j] {
				//TODO: Myhill Nerode theorem step 2
			}
		}
	}
}
