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

import "sort"

// Looks for all values in Y[] where x ^ Y[i] > Y[i] ^ x
func countNumberOfPairs(x int, y []int, lenY int, numberOfY []int) int {
	// If x is 0, then there cannot be any value in Y such that x^Y[i] > Y[i]^x
	if x == 0 {
		return 0
	}

	// If x is 1, then the number of pairs is equal to number of zeroes in Y[]
	if x == 1 {
		return numberOfY[0]
	}

	// Find number of elements in Y[] with values greater than x
	count := 0
	for i := 0; i < lenY; i++ {
		if y[i] > x {
			count = lenY - i
			break
		}
	}

	// If we have reached here, then x must be greater than 1, increase number of pairs for y=0 and y=1
	count += (numberOfY[0] + numberOfY[1])

	// Decrease number of pairs for x=2 and (y=4 or y=3)
	if x == 2 {
		count -= (numberOfY[3] + numberOfY[4])
	}

	// Increase number of pairs for x=3 and y=2
	if x == 3 {
		count += numberOfY[2]
	}

	return count
}

func numberOfPairs(x, y []int) int {
	lenX := len(x)
	lenY := len(y)
	if lenX <= 0 || lenY <= 0 {
		return 0
	}

	sort.Ints(y)

	//                 0  1  2  3  4
	numberOfY := []int{0, 0, 0, 0, 0}
	for i := 0; i < lenY; i++ {
		if y[i] < 5 {
			numberOfY[y[i]]++
		} else {
			break
		}
	}

	result := 0
	for i := 0; i < lenX; i++ {
		result += countNumberOfPairs(x[i], y, lenY, numberOfY)
	}

	return result
}
