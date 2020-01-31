// Count Triplets such that one of the numbers can be written as sum of the other two

// Given an array A[] of N integers. The task is to find the number of triples (i, j, k),
// where i, j, k are indices and (1 <= i < j < k <= N), such that in the set { A_i, A_j, A_k}
// at least one of the numbers can be written as the sum of the other two.

// Example 1:
// - Input: A[] = {1, 2, 3, 4, 5}
// - Output: 4
// - The valid triplets are: (1, 2, 3), (1, 3, 4), (1, 4, 5), (2, 3, 5)

// Example 2:
// - Input : A[] = {1, 1, 1, 2, 2}
// - Output : 6

package aarrays

import "testing"

type testcaseCountTheTriplets struct {
	in  []int
	out int
}

func doTestCountTheTriplets(t *testing.T, f func([]int) int) {
	cases := []testcaseCountTheTriplets{
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
	doTestCountTheTriplets(t, countTheTripletsGood)
}

// goos: darwin
// goarch: amd64
// BenchmarkCountTheTripletsBad-4    	     460	   2603353 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCountTheTripletsGood-4   	    1522	    794790 ns/op	   11121 B/op	      23 allocs/op

func BenchmarkCountTheTripletsBad(b *testing.B) {
	a := make([]int, 200)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	for i := 0; i < b.N; i++ {
		countTheTripletsBad(a)
	}
}

func BenchmarkCountTheTripletsGood(b *testing.B) {
	a := make([]int, 200)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	for i := 0; i < b.N; i++ {
		countTheTripletsGood(a)
	}
}
