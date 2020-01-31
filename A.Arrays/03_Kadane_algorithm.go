// Largest Sum Contiguous Subarray

// Write an efficient program to find the sum of contiguous subarray within
// a one-dimensional array of numbers which has the largest sum.
// Example:
// - Input: -2, -3, 4, -1, -2, 1, 5, -3
// - Output: 7
// (The subarray is 4, -1, -2, 1, 5)

package aarrays

import util "github.com/doanvanvinhtho/programming-interview-questions/Z.Util"

func kadaneAlgorithm(a []int) int {
	len := len(a)
	if len <= 0 {
		return 0
	}

	maxSoFar := a[0]
	maxEndingHere := a[0]
	for i := 1; i < len; i++ {
		maxEndingHere = util.IntMax(a[i]+maxEndingHere, a[i])
		maxSoFar = util.IntMax(maxEndingHere, maxSoFar)
	}

	return maxSoFar
}
