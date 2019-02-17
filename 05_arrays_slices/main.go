package main

import "fmt"

func main() {
	fmt.Println("Arrays And Slices")
	// Arrays
	// var fruitArr [2]string
	// // Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "orange"

	// Declare and Assign

	// fruitArr := [2]string{"apple", "orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])

	//Slices
	fruitSlice := []string{"apple", "orange", "Grape"}
	// fruitSlice[3] = "download"
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
	fmt.Println(cap(fruitSlice))
}
