package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	incrementer := 0
	counter := 100
	wg.Add(counter)

	fmt.Println("Start GORoutines:", runtime.NumGoroutine())
	for i := 0; i <= counter; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("The Value of incrementer is:", incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value:", incrementer)
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
