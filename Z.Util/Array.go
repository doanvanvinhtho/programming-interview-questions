package util

import "errors"

// ArrayEqual helps to compare 2 array/slice
func ArrayEqual(a, b []int) bool {
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

// ArrayMax helps to find maximum number of an array/slice
func ArrayMax(a []int) (int, error) {
	len := len(a)

	if len <= 0 {
		return -1, errors.New("The array is empty")
	}
	max := a[0]

	for i := 1; i < len; i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max, nil
}
