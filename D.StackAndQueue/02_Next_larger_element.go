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

func nextLargerElement1(a []int) []int {
	n := len(a)
	var result = make([]int, n, n)

	for i := 0; i < n; i++ {
		result[i] = -1
		for j := i + 1; j < n; j++ {
			if a[j] > a[i] {
				result[i] = a[j]
				break
			}
		}
	}

	return result
}

func nextLargerElement2(a []int) []int {
	n := len(a)
	var result = make([]int, n, n)

	var stack = make([]int, n, n)
	stackItems := -1

	for i := 0; i < n; i++ {
		if stackItems < 0 { // If stack is empty
			stackItems++
			stack[stackItems] = i // Stack.Push
			continue
		}
		for stackItems >= 0 {
			top := stack[stackItems] // Stack.Top
			if a[top] < a[i] {
				result[top] = a[i]
				stackItems-- // Stack.Pop
			} else {
				break
			}
		}
		stackItems++
		stack[stackItems] = i // Stack.Push
	}

	for stackItems >= 0 {
		top := stack[stackItems] // Stack.Top
		result[top] = -1
		stackItems-- // Stack.Pop
	}
	return result
}
