// Find the Missing Number

// You are given a list of n-1 integers and these integers are in the range of 1 to n.
// There are no duplicates in the list. One of the integers is missing in the list.
// Write an efficient code to find the missing integer.
// Example:
// Input: arr[] = {1, 2, 4, 6, 3, 7, 8}
// Output: 5

package aarrays

func missingNumberInArray(a []int, n int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return n*(n+1)/2 - sum
}
