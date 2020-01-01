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
