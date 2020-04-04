// Find number of pairs (x, y) in an array such that x^y > y^x

// Given two arrays X[] and Y[] of positive integers, find number of pairs
// such that x^y > y^x where x is an element from X[] and y is an element from Y[].

// Examples:
// Input: X[] = {2, 1, 6}, Y = {1, 5}
// Output: 3
// Explanation: There are total 3 pairs where pow(x, y) is greater than pow(y, x)
// Pairs are (2, 1), (2, 5) and (6, 1)
//
// Input: X[] = {10, 19, 18}, Y[] = {11, 15, 9}
// Output: 2
// Explanation: There are total 2 pairs where pow(x, y) is greater
// than pow(y, x) Pairs are (10, 11) and (10, 15)

package aarrays

import "testing"

type testcaseNumberOfPairs struct {
	inX []int
	inY []int
	out int
}

func TestNumberOfPairs(t *testing.T) {
	cases := []testcaseNumberOfPairs{
		{
			inX: []int{2, 1, 6},
			inY: []int{1, 5},
			out: 3,
		},
		{
			inX: []int{10, 19, 18},
			inY: []int{11, 15, 9},
			out: 2,
		},
		{
			inX: []int{},
			inY: []int{},
			out: 0,
		},
		{
			inX: []int{1},
			inY: []int{},
			out: 0,
		},
		{
			inX: []int{},
			inY: []int{1},
			out: 0,
		},
	}

	for i, c := range cases {
		r := numberOfPairs(c.inX, c.inY)
		if r != c.out {
			t.Errorf("Wrong %v got %v", i, r)
		}
	}
}
