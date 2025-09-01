package main

import "fmt"

func loops() {
	var arr1 = [...]string{"Mango", "Banana"}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("Index", i, "Fruit:", arr1[i])
	}
}
