// Rearrange an array in maximum minimum form

// Given a sorted array of positive integers, rearrange the array alternately
// i.e first element should be maximum value, second minimum value,
// third second max, fourth second min and so on.

// Examples:
// - Input : arr[] = {1, 2, 3, 4, 5, 6, 7}
// - Output: arr[] = {7, 1, 6, 2, 5, 3, 4}

package aarrays

func rearrangeArrayAlternatively(a []int) []int {
	len := len(a)

	var result = make([]int, len, len)

	minIndex := 0
	maxIndex := len - 1
	resultIndex := 0

	for minIndex < maxIndex {
		result[resultIndex] = a[maxIndex]
		maxIndex--
		resultIndex++

		result[resultIndex] = a[minIndex]
		minIndex++
		resultIndex++
	}
	if minIndex >= 0 && minIndex == maxIndex && maxIndex < len {
		result[resultIndex] = a[minIndex]
	}

	return result
}
