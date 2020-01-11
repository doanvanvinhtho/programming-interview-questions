package aarrays

import "testing"

func doTest(t *testing.T, subarray func(a []int, sum int) (int, int)) {
	startIndex, endIndex := subarray([]int{1, 4, 20, 3, 10, 5}, 33)
	if startIndex != 2 || endIndex != 4 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}

	startIndex, endIndex = subarray([]int{1, 4, 0, 0, 3, 10, 5}, 7)
	if startIndex != 1 || endIndex != 4 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}

	startIndex, endIndex = subarray([]int{1, 4, 0, 0, 3, 10, 5}, 15)
	if startIndex != 5 || endIndex != 6 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}

	startIndex, endIndex = subarray([]int{1, 4}, 0)
	if startIndex != -1 || endIndex != -1 {
		t.Errorf("Wrong, startIndex=%v, endIndex=%v", startIndex, endIndex)
	}
}

func TestSubarrayWithGivenSum(t *testing.T) {
	doTest(t, subarrayWithGivenSum)
	doTest(t, subarrayWithGivenSumNotGood)
}

// goos: darwin
// goarch: amd64
// BenchmarkSubarrayWithGivenSum-4          	24372546	        47.5 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSubarrayWithGivenSumNotGood-4   	10039954	       122 ns/op	       0 B/op	       0 allocs/op

func BenchmarkSubarrayWithGivenSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subarrayWithGivenSum([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 20, 3, 10, 5}, 33)
	}
}

func BenchmarkSubarrayWithGivenSumNotGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subarrayWithGivenSumNotGood([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 20, 3, 10, 5}, 33)
	}
}
