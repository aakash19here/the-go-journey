package main

import "fmt"

// Input: Array = [2, 1, 5, 1, 3, 2], k = 3
// Output: 9, because [5, 1, 3] has the maximum sum.

func maxSumSubarray(arr []int, k int) int {
	if k <= 0 || k > len(arr) {
		return 0
	}

	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	for i := k; i < len(arr); i++ {
		// old sum + new number entering - old number leaving
		windowSum = windowSum + arr[i] - arr[i-k]

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3

	fmt.Println(maxSumSubarray(arr, k))
}
