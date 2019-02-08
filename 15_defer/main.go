package main

import (
	"fmt"
)

func do() {
	n := 1
	defer fmt.Println("First", n)
	n = 2
	defer fmt.Println("Second", n)
	n = 3
	defer fmt.Println("third", n)

	var err error
	defer func() {
		if err != nil {
			fmt.Println("Error Happend", err)
		}
	}()
	err = fmt.Errorf("Happend")
}

func main() {
	do()
}
