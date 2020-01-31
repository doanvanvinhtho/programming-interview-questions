// Find the Missing Number

// You are given a list of n-1 integers and these integers are in the range of 1 to n.
// There are no duplicates in the list. One of the integers is missing in the list.
// Write an efficient code to find the missing integer.
// Example:
// Input: arr[] = {1, 2, 4, 6, 3, 7, 8}
// Output: 5

package aarrays

import "testing"

func TestMissingNumberInArray(t *testing.T) {
	if v := missingNumberInArray([]int{1, 2, 4, 6, 3, 7, 8}, 8); v != 5 {
		t.Errorf("Wrong, the missing number should be 5, but we get %v", v)
	}

	if v := missingNumberInArray([]int{1, 2, 4, 6, 3, 7, 5}, 8); v != 8 {
		t.Errorf("Wrong, the missing number should be 8, but we get %v", v)
	}
}
