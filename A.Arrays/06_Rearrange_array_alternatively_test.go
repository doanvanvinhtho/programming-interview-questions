// Rearrange an array in maximum minimum form

// Given a sorted array of positive integers, rearrange the array alternately
// i.e first element should be maximum value, second minimum value,
// third second max, fourth second min and so on.

// Example:
// - Input : arr[] = {1, 2, 3, 4, 5, 6, 7}
// - Output: arr[] = {7, 1, 6, 2, 5, 3, 4}

package aarrays

import (
	"testing"

	util "github.com/doanvanvinhtho/programming-interview-questions/Z.Util"
)

func TestRearrangeArrayAlternatively(t *testing.T) {
	result := rearrangeArrayAlternatively([]int{1, 2, 3, 4, 5, 6, 7})
	if util.ArrayIntEqual(result, []int{7, 1, 6, 2, 5, 3, 4}) == false {
		t.Errorf("Wrong 1: %v", result)
	}

	result = rearrangeArrayAlternatively([]int{1, 2, 3, 4, 5, 6})
	if util.ArrayIntEqual(result, []int{6, 1, 5, 2, 4, 3}) == false {
		t.Errorf("Wrong 2: %v", result)
	}

	result = rearrangeArrayAlternatively([]int{1})
	if util.ArrayIntEqual(result, []int{1}) == false {
		t.Errorf("Wrong 3: %v", result)
	}

	result = rearrangeArrayAlternatively([]int{})
	if util.ArrayIntEqual(result, []int{}) == false {
		t.Errorf("Wrong 4: %v", result)
	}
}
