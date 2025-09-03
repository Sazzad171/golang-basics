package main

import "fmt"

func functions() {
	personInfo("John", 18)
	fmt.Println("Addition result", addition(100, 50))
}

func personInfo(name string, age int) {
	fmt.Println("Name:", name, "Age:", age)
}

func addition(x int, y int) int {
	return x + y
}
