package main

import "fmt"

type Person struct {
	name string
	age  int
	job  string
}

func structures() {
	// multiple values, different data types
	var per1 Person
	var per2 Person

	per1.name = "John"
	per1.age = 20
	per1.job = "Student"

	per2.name = "Doe"
	per2.age = 40
	per2.job = "Engineer"

	fmt.Println(per1)
	fmt.Println(per2)

}
