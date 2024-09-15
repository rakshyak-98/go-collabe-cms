package main

import "fmt"

func main() {
	//Basic syntax
	// if x < 0 {
	//print or return
	// }

	//Like for, the if statement can start with a short statement to execute before the condition.
	if i := 10; i == 10 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
		//can not use variable i here
	}

	//Switch case

	c := 10
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Special character")
	case '1', '2':
		fmt.Println("Number")
	}
	fmt.Println("Unknown")
}
