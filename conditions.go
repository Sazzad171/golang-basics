package main

import "fmt"

func conditions() {
	var age int = 6

	if age < 18 {
		fmt.Println("You are not adult!")
	} else {
		fmt.Println("You are adult!")
	}

	// switch
	switch age {
	case 6:
		fmt.Println("You are kids!")
	case 18:
		fmt.Println("You are matured!")
	}
}
