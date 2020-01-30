package aarrays

import "testing"

type testcaseSubarrayWithGivenSum struct {
	inArray  []int
	inSum    int
	outStart int
	outEnd   int
}

func doTestSubarrayWithGivenSum(t *testing.T, subarray func(a []int, sum int) (int, int)) {
	cases := []testcaseSubarrayWithGivenSum{
		{
			inArray:  []int{1, 4, 20, 3, 10, 5},
			inSum:    33,
			outStart: 2,
			outEnd:   4,
		},
		{
			inArray:  []int{1, 4, 0, 0, 3, 10, 5},
			inSum:    7,
			outStart: 1,
			outEnd:   4,
		},
		{
			inArray:  []int{1, 4, 0, 0, 3, 10, 5},
			inSum:    15,
			outStart: 5,
			outEnd:   6,
		},
		{
			inArray:  []int{1, 4},
			inSum:    0,
			outStart: -1,
			outEnd:   -1,
		},
	}

	for i, v := range cases {
		startIndex, endIndex := subarray(v.inArray, v.inSum)
		if startIndex != v.outStart || endIndex != v.outEnd {
			t.Errorf("Wrong %v, startIndex=%v, endIndex=%v", i, startIndex, endIndex)
		}
	}
}

func TestSubarrayWithGivenSum(t *testing.T) {
	doTestSubarrayWithGivenSum(t, subarrayWithGivenSum)
	doTestSubarrayWithGivenSum(t, subarrayWithGivenSumNotGood)
}

// goos: darwin
// goarch: amd64
// BenchmarkSubarrayWithGivenSum-4          	24386794	        42.4 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSubarrayWithGivenSumNotGood-4   	12325627	       103 ns/op	       0 B/op	       0 allocs/op

func BenchmarkSubarrayWithGivenSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subarrayWithGivenSum([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 20, 3, 10, 5}, 33)
	}
}

func BenchmarkSubarrayWithGivenSumNotGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subarrayWithGivenSumNotGood([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 20, 3, 10, 5}, 33)
	}
}
