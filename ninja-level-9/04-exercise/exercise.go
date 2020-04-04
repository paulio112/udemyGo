package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	incrementer := 0
	counter := 100

	wg.Add(counter)

	var m sync.Mutex

	for i := 0; i < counter; i++ {
		go func() {

			m.Lock()
			v := incrementer
			v++
			incrementer = v

			fmt.Println("The Value of incrementer is:", incrementer)

			m.Unlock()

			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("End Value:", incrementer)
}

// Fix the above race condition using Mutex
// Hands-on exercise #4
// Fix the race condition you created in the previous exercise by using a mutex
// it makes sense to remove runtime.Gosched()
// code: https://github.com/GoesToEleven/go-programming
// video: 151

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
