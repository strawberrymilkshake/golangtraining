package main

import "fmt"

func action() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	c := action()
	for r := range c {
		fmt.Println(r)
	}
}
