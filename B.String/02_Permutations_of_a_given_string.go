// Write a program to print all permutations of a given string

// A permutation, also called an “arrangement number” or “order,” is a rearrangement of the elements
// of an ordered list S into a one-to-one correspondence with S itself. A string of length n has n! permutation.
// Source: Mathword(http://mathworld.wolfram.com/Permutation.html)
// Below are the permutations of string ABC.
// ABC ACB BAC BCA CBA CAB

package bstring

type permutationOfAStringOutput struct {
	result []string
}

func doPermutationOfAString(s string, left int, right int, o *permutationOfAStringOutput) {
	if left == right {
		o.result = append(o.result, s)
	} else {
		char := []rune(s)
		for i := left; i <= right; i++ {
			// Swap left, i
			char[left], char[i] = char[i], char[left]

			doPermutationOfAString(string(char), left+1, right, o)

			// Backtrack
			char[left], char[i] = char[i], char[left]
		}
	}
}

func permutationOfAString(s string) []string {
	o := permutationOfAStringOutput{
		result: []string{},
	}

	doPermutationOfAString(s, 0, len(s)-1, &o)

	return o.result
}
