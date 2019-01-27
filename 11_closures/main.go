package main

import "fmt"

func addr() func(int) int { // what the hell, function that return func that contains function
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := addr()

	for i := 0; i < 10; i++ {
		fmt.Println(i, sum(i))
	}
}
