package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(*b)
	fmt.Println(*&a)

	// Change Value With Ponters
	*b = 10
	fmt.Println(a)
}
