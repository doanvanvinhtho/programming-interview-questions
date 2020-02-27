// Merge two sorted arrays
// Given two sorted arrays, the task is to merge them in a sorted manner.

// Example 1:
// - Input: arr1[] = { 1, 3, 4, 5}, arr2[] = {2, 4, 6, 8}
// - Output: arr3[] = {1, 2, 3, 4, 4, 5, 6, 8}

// Example 2:
// - Input: arr1[] = { 5, 8, 9}, arr2[] = {4, 7, 8}
// - Output: arr3[] = {4, 5, 7, 8, 8, 9}

package aarrays

func mergeTwoSortedArrays(a, b []int) []int {
	lenA := len(a)
	lenB := len(b)

	if lenA <= 0 && lenB <= 0 {
		return []int{}
	}

	r := make([]int, lenA+lenB, lenA+lenB)
	iA, iB, iR := 0, 0, 0

	for iA < lenA && iB < lenB {
		if a[iA] <= b[iB] {
			r[iR] = a[iA]
			iA++
		} else {
			r[iR] = b[iB]
			iB++
		}
		iR++
	}

	for ; iA < lenA; iA++ {
		r[iR] = a[iA]
		iR++
	}

	for ; iB < lenB; iB++ {
		r[iR] = b[iB]
		iR++
	}

	return r
}
