package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup
	var counter int64
	wg.Add(2) //number of wg.Done() needed to proceed with wg.Wait()
	go func() {
		for i := 0; i < 40; i++ {
			time.Sleep(time.Duration(9 * time.Millisecond))
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter incrementer by f1:", counter)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 40; i++ {
			time.Sleep(time.Duration(21 * time.Millisecond))
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter incremented by f2:", counter)
		}
		wg.Done()
	}()
	wg.Wait()
}
