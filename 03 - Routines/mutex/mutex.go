package main
import (
  "fmt"
"runtime"
"sync"
"time"
)
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup
  var mutex sync.Mutex
  var counter int
	wg.Add(2) //number of wg.Done() needed to proceed with wg.Wait()
	go func() {
		for i := 0; i < 450; i++ {
      			time.Sleep(time.Duration(9 * time.Millisecond))
      mutex.Lock()
      counter++
			fmt.Println("Counter incrementer by f1:", counter)

      mutex.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 450; i++ {
      time.Sleep(time.Duration(21 * time.Millisecond))
      mutex.Lock()
      counter++
			fmt.Println("Counter incremented by f2:", counter)
      mutex.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}
