/*
In lieu of adding exception handlers, the Go creators exploited Go’s ability to return multiple values. The most commonly used Go technique for issuing errors is to return the error as the last value in a return.

A panic typically means something went unexpectedly wrong. Mostly used to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

Panic recovery in Go depends on a feature of the language called deferred functions. Go has the ability to guarantee the execution of a function at the moment its parent function returns. This happens regardless of whether the reason for the parent function’s return is a return statement, the end of the function block, or a panic.
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("go.mod")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f)
}

// Resource 1: https://go.dev/blog/error-handling-and-go
// Resource 1: https://earthly.dev/blog/golang-errors/
// Resource 1: https://go.dev/blog/defer-panic-and-recover
