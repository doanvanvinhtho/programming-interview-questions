// Write a program to print all permutations of a given string

// A permutation, also called an “arrangement number” or “order,” is a rearrangement of the elements
// of an ordered list S into a one-to-one correspondence with S itself. A string of length n has n! permutation.
// Source: Mathword(http://mathworld.wolfram.com/Permutation.html)
// Below are the permutations of string ABC.
// ABC ACB BAC BCA CBA CAB

package bstring

import (
	"testing"

	util "github.com/doanvanvinhtho/programming-interview-questions/Z.Util"
)

type testcasePermutation struct {
	inS string
	out []string
}

func TestPermutationOfAString(t *testing.T) {
	cases := []testcasePermutation{
		{
			inS: "ABC",
			out: []string{"ABC", "ACB", "BAC", "BCA", "CBA", "CAB"},
		},
		{
			inS: "ab",
			out: []string{"ab", "ba"},
		},
		{
			inS: "z",
			out: []string{"z"},
		},
	}
	for i, c := range cases {
		r := permutationOfAString(c.inS)
		if util.ArrayStringEqual(r, c.out) == false {
			t.Errorf("Wrong %v got %v", i, r)
		}
	}
}
