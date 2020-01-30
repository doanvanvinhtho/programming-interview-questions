// Count Triplets such that one of the numbers can be written as sum of the other two

// Given an array A[] of N integers. The task is to find the number of triples (i, j, k),
// where i, j, k are indices and (1 <= i < j < k <= N), such that in the set { A_i, A_j, A_k}
// at least one of the numbers can be written as the sum of the other two.

// Examples 1:
// - Input: A[] = {1, 2, 3, 4, 5}
// - Output: 4
// - The valid triplets are: (1, 2, 3), (1, 3, 4), (1, 4, 5), (2, 3, 5)

// Examples 2:
// - Input : A[] = {1, 1, 1, 2, 2}
// - Output : 6

package aarrays

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

// func countTheTripletsGood(a []int) int {
// 	count := 0
// 	len := len(a)
// 	return count
// }
