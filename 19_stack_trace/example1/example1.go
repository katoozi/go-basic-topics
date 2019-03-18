// Sample program to show how to read stack trace

package main

func main() {
	example(make([]string, 2, 4), "hello", 10)
}

func example(slice []string, str string, i int) {
	panic("Want stack trace")
}
