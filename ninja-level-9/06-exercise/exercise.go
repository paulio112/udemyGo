package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}

// Create a program that prints out your OS and ARCH. Use the following commands to run it
// go run
// go build
// go install
// code: https://github.com/GoesToEleven/go-programming
// video: 153
