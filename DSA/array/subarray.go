package array

import "fmt"

/*
Q: https://leetcode.com/problems/maximum-average-subarray-i/
*/

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5
	result := make([]float64, 9)

	windowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd += 1 {
		windowSum += arr[windowEnd]
		if windowEnd >= k-1 {
			result = append(result, float64(windowSum)/float64(k))
			windowSum -= arr[windowStart]
			windowStart += 1
		}
	}
	fmt.Println(result)
}
