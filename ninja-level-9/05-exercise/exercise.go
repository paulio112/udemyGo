package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int64

	const gr = 100
	wg.Add(gr)
	fmt.Println("Start GORoutines:", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter is:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End  GORoutines:", runtime.NumGoroutine())

}

// Using goroutines, create an incrementer program
// have a variable to hold the incrementer value
// launch a bunchclear

// each goroutine should
// read the incrementer value
// store it in a new variable
// yield the processor with runtime.Gosched()
// increment the new variable
// write the value in the new variable back to the incrementer variable
// use waitgrcleoups to wait for all of your goroutines to finish
// the above will create a race condition.
// Prove that it is a race condition by using the -race flag
// if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
// code: https://github.com/GoesToEleven/go-programming
// video: 150
