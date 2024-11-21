package array

import (
	"math"
)

// Q: Given an array of positive numbers and a positive number 'S', find the length of the smallest
// contiguous subarray whose sum is greater than or equal to 'S'. Return 0, if no such subarray exists.

// input : [2,1,5,2,3,2]
// input : sum = 7

// output : length of smallest sub array with sum = 7
// i = 0, 2 ->  sum 2
// i = 1, 2,1 -> sum = 3
// i = 2, 2,1,5 -> sum = 8 // may be answer smallest length 3
// i = 3, 1,5,2 -> sum = 8
// i = 4, 5,2 -> sum = 7 // may be answer smallest length 2
// i = 5,  2, 3, 2 -> sum = 7 // not smallest length
// i = 6, 3,2 -> sum = 5
// i = 7, 2 -> sum = 2

func find_smallest_substring(s []int, sum int) int {
	windowSum := 0
	windowStart := 0
	min_length := math.MaxInt

	for windowEnd := 0; windowEnd < len(s); windowEnd += 1 {
		windowSum += s[windowEnd]
		for windowSum >= sum {
			min_length = int(math.Min(float64(min_length), float64(windowEnd-windowStart+1)))
			windowSum -= s[windowStart]
			windowStart += 1
		}
	}
	if min_length == math.MaxInt {
		return 0
	}
	return min_length
}
