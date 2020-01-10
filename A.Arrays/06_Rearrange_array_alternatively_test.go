package aarrays

import "testing"

func TestRearrangeArrayAlternatively(t *testing.T) {
	result := rearrangeArrayAlternatively([]int{1, 2, 3, 4, 5, 6, 7})
	if equal(result, []int{7, 1, 6, 2, 5, 3, 4}) == false {
		t.Errorf("Wrong 1: %v", result)
	}

	result = rearrangeArrayAlternatively([]int{1, 2, 3, 4, 5, 6})
	if equal(result, []int{6, 1, 5, 2, 4, 3}) == false {
		t.Errorf("Wrong 2: %v", result)
	}

	result = rearrangeArrayAlternatively([]int{1})
	if equal(result, []int{1}) == false {
		t.Errorf("Wrong 3: %v", result)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
