package main

import "fmt"

func main() {
	fmt.Println("Conditionals")
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less then %d\n", x, y)
	} else {
		fmt.Printf("%d is grather then %d\n", x, y)
	}

	color := "red"
	if color == "red" { // &&(and) and ||(or)
		fmt.Println("Color is Red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}

	// Switch case statement

	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}
}
