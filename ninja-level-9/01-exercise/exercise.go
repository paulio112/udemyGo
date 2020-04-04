package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin GORoutine:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("First Function Running")
		wg.Done()
	}()

	go func() {
		fmt.Println("First Function Running")
		wg.Done()
	}()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid GORoutine:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End GORoutine:", runtime.NumGoroutine())

}

// Hands-on exercise #1
// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists
// code: https://github.com/GoesToEleven/go-programming
// video: 148
