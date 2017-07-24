package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	in := make(chan bool)
	out := make(chan string)
	wg.Add(2)
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
		wg.Done()
	}()
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Duration(13 * time.Millisecond))
			in <- false
			fmt.Println(<-out)
		}
		wg.Done()
	}()
	wg.Wait()
}
