// Next Greater Element

// Given an array, print the Next Greater Element (NGE) for every element.
// The Next greater Element for an element x is the first greater element
// on the right side of x in array. Elements for which no greater element
// exist, consider next greater element as -1.

// Examples:
// - For any array, rightmost element always has next greater element as -1.
// - For an array which is sorted in decreasing order, all elements have next greater element as -1.
// - For the input array [4, 5, 2, 25], the next greater elements for each element are as follows.
//   Element       NGE
//    4      -->   5
//    5      -->   25
//    2      -->   25
//    25     -->   -1
//
// - For the input array [13, 7, 6, 12], the next greater elements for each element are as follows.
//   Element        NGE
//    13     -->    -1
//    7      -->    12
//    6      -->    12
//    12     -->    -1

package dstackandqueue

import (
	"testing"

	util "github.com/doanvanvinhtho/programming-interview-questions/Z.Util"
)

func doTest(t *testing.T, f func([]int) []int) {
	// For any array, rightmost element always has next greater element as -1.
	result := f([]int{1})
	if util.ArrayIntEqual(result, []int{-1}) == false {
		t.Errorf("Wrong 1: %v", result)
	}

	//For an array which is sorted in decreasing order, all elements have next greater element as -1.
	result = f([]int{3, 2, 1})
	if util.ArrayIntEqual(result, []int{-1, -1, -1}) == false {
		t.Errorf("Wrong 2: %v", result)
	}

	result = f([]int{4, 5, 2, 25})
	if util.ArrayIntEqual(result, []int{5, 25, 25, -1}) == false {
		t.Errorf("Wrong 3: %v", result)
	}

	result = f([]int{13, 7, 6, 12})
	if util.ArrayIntEqual(result, []int{-1, 12, 12, -1}) == false {
		t.Errorf("Wrong 4: %v", result)
	}
}

func TestNextLargerElement(t *testing.T) {
	doTest(t, nextLargerElement1)
	doTest(t, nextLargerElement2)
}

// goos: darwin
// goarch: amd64
// BenchmarkNextLargerElement1-4   	     913	   1284447 ns/op	   32855 B/op	       1 allocs/op
// BenchmarkNextLargerElement2-4   	   48710	     24579 ns/op	   65537 B/op	       2 allocs/op

func BenchmarkNextLargerElement1(b *testing.B) {
	a := []int{13, 7, 6, 12}
	for i := 0; i < 10; i++ {
		a = append(a, a...)
	}
	// 0   1   2   3    4    5     6     7     8     9
	// 8, 16, 32, 64, 128, 256, 1024, 2048, 4096, 8192

	for i := 0; i < b.N; i++ {
		nextLargerElement1(a)
	}
}

func BenchmarkNextLargerElement2(b *testing.B) {
	a := []int{13, 7, 6, 12}
	for i := 0; i < 10; i++ {
		a = append(a, a...)
	}
	// 0   1   2   3    4    5     6     7     8     9
	// 8, 16, 32, 64, 128, 256, 1024, 2048, 4096, 8192

	for i := 0; i < b.N; i++ {
		nextLargerElement2(a)
	}
}
