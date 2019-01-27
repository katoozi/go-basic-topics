package main

import "fmt"

func main() {
	fmt.Println("Range")

	ids := []int{33, 5215, 512, 515}

	// loop through ids
	for i, id := range ids {
		fmt.Println(i, id)
	}

	// Not using Index
	for _, id := range ids {
		fmt.Println(id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range With Map
	emails := map[string]string{"Bob": "bob@gmail.com", "Ali": "Ali@gmail.com"}

	for key, value := range emails {
		fmt.Println(key, value)
	}
}
