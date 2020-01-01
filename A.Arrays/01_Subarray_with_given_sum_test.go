package aarrays

import "testing"

func TestSubarrayWithGivenSum(t *testing.T) {
	startIndex, endIndex := subarrayWithGivenSum([]int{1, 4, 20, 3, 10, 5}, 33)
	if startIndex != 2 || endIndex != 4 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}

	startIndex, endIndex = subarrayWithGivenSum([]int{1, 4, 0, 0, 3, 10, 5}, 7)
	if startIndex != 1 || endIndex != 4 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}

	startIndex, endIndex = subarrayWithGivenSum([]int{1, 4, 0, 0, 3, 10, 5}, 15)
	if startIndex != 5 || endIndex != 6 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}

	startIndex, endIndex = subarrayWithGivenSum([]int{1, 4}, 0)
	if startIndex != -1 || endIndex != -1 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}
}
