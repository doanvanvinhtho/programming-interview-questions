package util

import "errors"

// ArrayIntEqual helps to compare 2 array/slice
func ArrayIntEqual(a, b []int) bool {
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

// ArrayStringEqual helps to compare 2 array/slice
func ArrayStringEqual(a, b []string) bool {
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

// ArrayIntMax helps to find maximum number of an array/slice
func ArrayIntMax(a []int) (int, error) {
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
