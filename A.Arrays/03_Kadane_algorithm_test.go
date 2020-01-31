// Largest Sum Contiguous Subarray

// Write an efficient program to find the sum of contiguous subarray within
// a one-dimensional array of numbers which has the largest sum.
// Example:
// - Input: -2, -3, 4, -1, -2, 1, 5, -3
// - Output: 7
// (The subarray is 4, -1, -2, 1, 5)

package aarrays

import "testing"

type testcaseKadaneAlgorithm struct {
	in  []int
	out int
}

func TestKadaneAlgorithm(t *testing.T) {
	cases := []testcaseKadaneAlgorithm{
		{
			in:  []int{-2, -3, 4, -1, -2, 1, 5, -3},
			out: 7,
		},
		{
			in:  []int{-1, -2},
			out: -1,
		},
		{
			in:  []int{-1, 2},
			out: 2,
		},
		{
			in:  []int{1, 2},
			out: 3,
		},
		{
			in:  []int{},
			out: 0,
		},
		{
			in:  []int{1},
			out: 1,
		},
	}

	for i, v := range cases {
		if out := kadaneAlgorithm(v.in); out != v.out {
			t.Errorf("Wrong %v, got=%v", i, out)
		}
	}
}
