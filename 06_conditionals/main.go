package main

import (
	"errors"
	"fmt"
)

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

	var (
		even   int
		odd    int
		zeroes int
		totla  int
	)

	numbers := []int{1, 2, 3, 4, 5, 0, -1, 8, 9, 10}

	var err error

Abort:
	for _, n := range numbers {
		switch {
		case n == 0:
			zeroes++
		case n%2 == 0:
			even++
		case n%2 == 1:
			odd++
		default:
			err = errors.New("Found Negetive Number")
			break Abort
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Even %d, Odd %d, Zeroes %d, total %d\n", even, odd, zeroes, totla)
}
