// Find subarray with given sum (Nonnegative Numbers)

// Given an unsorted array of nonnegative integers, find a continuous
// subarray which adds to a given number.
// Examples:
// Input: arr[] = {1, 4, 20, 3, 10, 5}, sum = 33
// Output: Sum found between indexes 2 and 4

package aarrays

func subarrayWithGivenSum(a []int, sum int) (int, int) {
	if len(a) <= 0 {
		return -1, -1
	}

	startIndex := 0
	currentSum := a[0]

	for i := 1; i <= len(a); i++ {
		for currentSum > sum && startIndex < i-1 {
			currentSum -= a[startIndex]
			startIndex++
		}

		if currentSum == sum {
			return startIndex, i - 1
		}

		if i < len(a) {
			currentSum += a[i]
		}
	}

	return -1, -1
}

func subarrayWithGivenSumNotGood(a []int, sum int) (int, int) {
	for startIndex := 0; startIndex < len(a); startIndex++ {
		currentSum := 0
		for endIndex := startIndex; endIndex < len(a); endIndex++ {
			currentSum += a[endIndex]
			if currentSum == sum {
				return startIndex, endIndex
			}
		}
	}

	return -1, -1
}
