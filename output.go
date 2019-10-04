package main

import "fmt"

func printTransitionTable() {
	fmt.Printf("\t")
	for i := range inputSymbols {
		fmt.Printf("%s\t", i)
	}
	fmt.Printf("\n")
	for i := range states {
		fmt.Printf(states[i].name)
		for t := range states[i].transitions {
			fmt.Printf("\t%s", states[i].transitions[t].name)
		}
		fmt.Printf("\n")
	}
}
