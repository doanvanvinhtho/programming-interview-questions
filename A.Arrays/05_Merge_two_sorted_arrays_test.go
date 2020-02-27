// Merge two sorted arrays
// Given two sorted arrays, the task is to merge them in a sorted manner.

// Example 1:
// - Input: arr1[] = {1, 3, 4, 5}, arr2[] = {2, 4, 6, 8}
// - Output: arr3[] = {1, 2, 3, 4, 4, 5, 6, 8}

// Example 2:
// - Input: arr1[] = {5, 8, 9}, arr2[] = {4, 7, 8}
// - Output: arr3[] = {4, 5, 7, 8, 8, 9}

package aarrays

import (
	"testing"

	util "github.com/doanvanvinhtho/programming-interview-questions/Z.Util"
)

type testcaseMergeTwoSortedArrays struct {
	inA []int
	inB []int
	out []int
}

func TestMergeTwoSortedArrays(t *testing.T) {
	cases := []testcaseMergeTwoSortedArrays{
		{
			inA: []int{1, 3, 4, 5},
			inB: []int{2, 4, 6, 8},
			out: []int{1, 2, 3, 4, 4, 5, 6, 8},
		},
		{
			inA: []int{5, 8, 9},
			inB: []int{4, 7, 8},
			out: []int{4, 5, 7, 8, 8, 9},
		},
		{
			inA: []int{},
			inB: []int{},
			out: []int{},
		},
		{
			inA: []int{},
			inB: []int{4, 7, 8},
			out: []int{4, 7, 8},
		},
		{
			inA: []int{5, 8, 9},
			inB: []int{},
			out: []int{5, 8, 9},
		},
	}

	for i, c := range cases {
		r := mergeTwoSortedArrays(c.inA, c.inB)
		if util.ArrayEqual(r, c.out) == false {
			t.Errorf("Wrong %v got %v", i, r)
		}
	}
}
