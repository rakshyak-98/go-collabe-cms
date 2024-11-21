package array

import (
	"fmt"
	"math"
)

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

/*
Q: https://leetcode.com/problems/largest-subarray-length-k/
 - Given an array of positive numbers and a positive number ‘k’ , find the maximum sum of any contiguous subarray of size ‘k’
*/

func maxSubArrayOfSizeK(arr []float64, K int) float64 {
	var maxSum float64 = math.MinInt
	var windowSum float64 = 0.0
	var windowStart int = 0

	for i := 0; i < len(arr); i += 1 {
		windowSum += arr[i]
		if i >= K-1 {
			maxSum = math.Max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart += 1
		}
	}
	return maxSum
}
