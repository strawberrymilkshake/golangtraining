package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) //number of wg.Done() needed to proceed with wg.Wait()
	go func() {
		for i := 0; i < 45; i++ {
			fmt.Println("Sic!")
			//time.Sleep(time.Duration(21 * time.Millisecond))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 45; i++ {
			fmt.Println("Arrr!")
			//time.Sleep(time.Duration(10 * time.Millisecond))
		}
		wg.Done()
	}()
	wg.Wait()
}
