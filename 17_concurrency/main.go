package main

import "fmt"

// Counter is simple struct for manage goroutine
type Counter struct {
	c    chan int
	done chan struct{}
	i    int
}

// GetSource will return counter c element
func (c *Counter) GetSource() <-chan int {
	return c.c
}

// Stop will stop the goroutine
func (c *Counter) Stop() {
	c.done <- struct{}{}
}

// NewCounter will return a new Counter struct
func NewCounter() *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.done = make(chan struct{})

	go func() {
		for {
			select {
			case counter.c <- counter.i:
				counter.i++
			case <-counter.done:
				return
			}
		}
	}()

	return counter
}

func main() {
	c := NewCounter()
	read := c.GetSource()
	fmt.Printf("%d\n", <-read) // wait for write in to c channel and then read it and go to next line
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	c.Stop()

	fmt.Printf("%d\n", <-read) // will rise the 'fatal error: all goroutines are asleep - deadlock!'

}
