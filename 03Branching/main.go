//Unlike other programing language Go don't have sepaRATE FOR, While and DO-WHITE
// Like a C for
// for init; condition; post { }

// Like a C while
// for condition { }

// Like a C for(;;)
// for { }

package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	//The init and post statements are optional.
	j := 1
	for j < 10 {
		j += j
	}
	fmt.Println(j)

	//Range: Used with loop to iterate over each element in arrays, strings and other data structures
	//It returns two values -> index and its element
	var pow = []int{1, 4, 8, 16, 32, 64}

	for key, value := range pow {
		fmt.Printf("%d contain %d\n", key, value)
	}

	// //We can skip the index or value by assigning to _.
	// for key := range pow {
	// 	//print key
	// }
	// for _, value := range pow {
	// 	//print value
	// }
}
