package aarrays

import "testing"

type testcase struct {
	in  []int
	out int
}

func doTestCountTheTriplets(t *testing.T, f func([]int) int) {
	cases := []testcase{
		{
			in:  []int{1, 2, 3, 4, 5},
			out: 4,
		},
		{
			in:  []int{1, 1, 1, 2, 2},
			out: 6,
		},
		{
			in:  []int{},
			out: 0,
		},
		{
			in:  []int{1},
			out: 0,
		},
		{
			in:  []int{1, 2},
			out: 0,
		},
		{
			in:  []int{1, 2, 3},
			out: 1,
		},
	}

	for i, v := range cases {
		if out := f(v.in); out != v.out {
			t.Errorf("Wrong %v, got=%v", i, out)
		}
	}
}

func TestCountTheTriplets(t *testing.T) {
	doTestCountTheTriplets(t, countTheTripletsBad)
	// doTestCountTheTriplets(t, countTheTripletsGood)
}
