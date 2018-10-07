package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr[0])
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Manggoes", "Coconut"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[2:3])
	fmt.Println(len(fruitSlice))
}
