/*
In lieu of adding exception handlers, the Go creators exploited Go’s ability to return multiple values. The most commonly used Go technique for issuing errors is to return the error as the last value in a return.

A panic typically means something went unexpectedly wrong. Mostly used to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

Panic recovery in Go depends on a feature of the language called deferred functions. Go has the ability to guarantee the execution of a function at the moment its parent function returns. This happens regardless of whether the reason for the parent function’s return is a return statement, the end of the function block, or a panic.
*/

/*
Defer: Used to ensure that a function call is performed later. The deferred function is executed after the surrounding function returnss. Executed in LIFO order.
Panic: Build-in function, that stop the normal execution of goroutine. Tipically used to signal unrecoverable error.
Recover: It is a build-in function which regain control of a panicking goroutine. It handle the panic and resume normal execution. Must be called in deferred function, otherwise it will not stop panic.
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

	// Defer Example
	deferEx()

	//Panic, Recover Example
	mightPanic()
	fmt.Println("Normal execution after panic")

}

func deferEx() {
	fmt.Println("Hello from Defer!")
	defer fmt.Println("Deferred: One")
	defer fmt.Println("Deferred: Two")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}

func mightPanic() {
	defer handlePanic() //comment this line for experience 'panic' 🙂
	panic("Something went wrong")
	fmt.Println("Hello From mightPani function") // will not execute this line
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recover from panic: ", r)
	}
}

// Code After Panic: Any code following a panic statement in the same function will not be executed.
// Deferred Functions: Deferred functions will still be executed even if a panic occurs, allowing you to perform cleanup or logging operations.
// Stack Unwinding: When a panic occurs, Go unwinds the stack, running deferred functions and propagating the panic up the call stack until it is recovered or the program terminates.

// Resource 1: https://go.dev/blog/error-handling-and-go
// Resource 1: https://earthly.dev/blog/golang-errors/
// Resource 1(Use of defer in Resource allocation/deallocation): https://go.dev/blog/defer-panic-and-recover
