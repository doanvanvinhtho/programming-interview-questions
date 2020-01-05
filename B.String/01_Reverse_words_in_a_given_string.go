// Reverse words in a given string

// Example: Let the input string be “i like this program very much”.
// The function should change the string to “much very program this like i”

package bstring

import "unicode"

func reverseWordsInString1(s string) string {
	result := []rune{}
	temp := []rune{}

	char := []rune(s)
	i := len(char) - 1

	for i >= 0 {
		if unicode.IsSpace(char[i]) {
			if len(result) > 0 {
				result = append(result, char[i])
			}
			result = append(result, temp...)
			temp = []rune{}
		} else {
			temp = append([]rune{char[i]}, temp...)
		}
		i--
	}
	if len(temp) > 0 {
		result = append(result, ' ')
		result = append(result, temp...)
	}

	return string(result)
}

// Algorithm:
// - Initially, reverse the individual words of the given string one by one,
//   for the above example, after reversing individual words the string should be
//   “i ekil siht margorp yrev hcum”.
//
// - Reverse the whole string from start to end to get the desired output
//   “much very program this like i” in the above example.
func reverseWordsInString2(s string) string {
	char := []rune(s)
	startIndex := -1

	for i := 0; i < len(char); i++ {
		if unicode.IsSpace(char[i]) {
			reverseWords2(char, startIndex+1, i-1)
			startIndex = i
		}
	}
	reverseWords2(char, startIndex+1, len(char)-1)
	reverseWords2(char, 0, len(char)-1)

	return string(char)
}

func reverseWords2(s []rune, startIndex int, endIndex int) {
	len := len(s)
	for (startIndex >= 0) && (startIndex < endIndex) && (endIndex <= len-1) {
		s[startIndex], s[endIndex] = s[endIndex], s[startIndex]
		startIndex++
		endIndex--
	}
}
