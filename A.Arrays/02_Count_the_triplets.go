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

import util "github.com/doanvanvinhtho/programming-interview-questions/Z.Util"

func countTheTripletsBad(a []int) int {
	count := 0
	len := len(a)

	for i := 0; i < len-2; i++ {
		for j := i + 1; j < len-1; j++ {
			for k := j + 1; k < len; k++ {
				if (a[i]+a[j] == a[k]) ||
					(a[i]+a[k] == a[j]) ||
					(a[j]+a[k] == a[i]) {
					count++
				}
			}
		}
	}

	return count
}

func countTheTripletsGood(a []int) int {
	max, e := util.ArrayMax(a)
	if e != nil {
		return 0
	}

	count := 0
	len := len(a)

	frequency := make(map[int]int)
	frequency[0] = 0
	frequency[max+1] = 0
	for i := 0; i < len; i++ {
		frequency[a[i]]++
	}

	// Case 1: 0, 0, 0
	count += (frequency[0] * (frequency[0] - 1) * (frequency[0] - 2)) / 6

	// Case 2: 0, x, x
	for i := 1; i <= max; i++ {
		count += (frequency[0] * frequency[i] * (frequency[i] - 1)) / 2
	}

	// Case 3: x, x, 2*x
	for i := 1; 2*i <= max; i++ {
		count += (frequency[i] * (frequency[i] - 1)) / 2 * frequency[2*i]
	}

	// Case 4: x, y, x + y
	// Iterate through all pairs (x, y)
	for i := 1; i <= max; i++ {
		for j := i + 1; i+j <= max; j++ {
			count += frequency[i] * frequency[j] * frequency[i+j]
		}
	}

	return count
}
