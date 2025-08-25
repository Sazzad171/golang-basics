package main

import "fmt"

var entryType string

func variables() {
	entryType = "Person"
	var personName string = "John"
	var personLastName = "Doe"
	age := 20

	fmt.Println(entryType)
	fmt.Println(personName)
	fmt.Println(personLastName)
	fmt.Printf("Age: %v and %T \n", age, age)

	const piValue float32 = 3.1416

	fmt.Print(piValue, "\n")

	/*
		variable types:
		string
		bool
		float32
		float64
		int
	*/
}
