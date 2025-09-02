package main

import "fmt"

func slices() {
	slice1 := []int{}
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := []string{"Go", "Slices"}
	fmt.Println(slice2)

	// slice from an array
	arr1 := [2]string{"mango", "pineapple"}
	sliceFromArray := arr1[0:2]
	fmt.Printf("Slice from array: %v \n", sliceFromArray)

	// slice with make
	sliceWithMake := make([]int, 2, 4)
	fmt.Printf("Slice with make function: %v \n", sliceWithMake)

	// slice append
	sliceAppend := []int{1, 2, 3}

	sliceAppend = append(sliceAppend, 4, 5)
	fmt.Printf("Slice after append %v \n", sliceAppend)

	sliceAppend2 := []int{11, 12, 13}
	sliceAppendNewVar := append(sliceAppend, sliceAppend2...)

	fmt.Printf("Slice append on new var %v \n", sliceAppendNewVar)
}
