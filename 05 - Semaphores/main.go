package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan bool)
	out := make(chan string)
	done := make(chan bool)
	defer func() {
		close(in)
		close(out)
		close(done)
	}()

	go func() {
		for c := range in {
			if c {
				out <- "Input is true"
			} else {
				out <- "Input is false"
			}
		}
	}()
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Duration(10 * time.Millisecond))
			in <- true
			fmt.Println(<-out)
		}
		done <- true
	}()
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Duration(13 * time.Millisecond))
			in <- false
			fmt.Println(<-out)
		}
		done <- true
	}()
	<-done
	<-done
}
