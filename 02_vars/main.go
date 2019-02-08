package main

import "fmt"

var name = "mohammad" // global var

func main() {
	// var name = "mohammad"
	var age int32 = 21
	const isCool = true
	family := "katoozi"

	size, wieght := 1.3, 105

	fmt.Println(name, age, isCool, family, size, wieght)
	fmt.Printf("%T\n", isCool)

	size = 1.2 // for reassign you have to use = not :=
}
