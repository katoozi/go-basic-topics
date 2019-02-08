package main

import "fmt"

func main() {
	// long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("FizzBuzz %d\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Fizz %d\n", i)
		} else if i%5 == 0 {
			fmt.Printf("Buzz %d\n", i)
		}
	}

	var (
		even  int
		odd   int
		total int
	)

	numbers := []int{1, 2, 3, 5, 0, 7, 8, 9, 10}
Abort:
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			total++
			if n == 0 {
				continue Abort
			}
			if n%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	fmt.Printf("Even %d, Odd %d, totoal %d\n", even, odd, total)
}
