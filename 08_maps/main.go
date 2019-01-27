package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Define a map
	// emails := make(map[string]string)
	emails := map[string]string{"Bob": "dksmkd@gmail.com", "Ali": "ali@gmail.com"}
	// Assign to map
	// emails["Bob"] = "asskdmm@gmail.com"
	// emails["Ali"] = "ali@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Ali"])
	fmt.Println(len(emails))

	// Delete From Map
	delete(emails, "Ali")
	fmt.Println(emails)
}
