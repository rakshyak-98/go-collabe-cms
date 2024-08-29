// One unusual feature in GO is that functions and methods can return multiple values

package main

import (
	"fmt"
)

//Syntax: func functionName(parameterList) (returnTypeList){ }

//Syntax: func functionName(parameterList) (returnTypeList){ }

// One return value
func add(a int, b int) int {
	return a + b
}

// Multiple return value
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values: When function value initialize to their zero value of type;
// The current values of the result parameters are used as the returned values.
func swap(a, b string) (f, s string) {
	f = b
	s = a
	return
}

func main() {
	sum := add(10, 2)
	fmt.Println("Sum: ", sum)

	res, err := divide(10.25, 0.25)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Division: ", res)
	}

	_, err = divide(10.0, 0.0)

	if err != nil {
		fmt.Println("Error:", err)
	}

	a, b := swap("Akash", "Tripathy")

	fmt.Println("Swapped: ", a, b)
}
