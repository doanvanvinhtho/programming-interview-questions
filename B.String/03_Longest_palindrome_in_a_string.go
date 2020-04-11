// Longest Palindromic Substring

// Given a string, find the longest substring which is palindrome.
// For example, if the given string is “forgeeksskeegfor”, the output should be “geeksskeeg”.

package bstring

func longestPalindromicSubstring1(s string) string {
	len := len(s)

	if len <= 0 {
		return ""
	}

	char := []rune(s)

	resultLeft := -1
	resultRight := -1
	resultLen := -1

	for left := 0; left < len; left++ {
		for right := len - 1; right >= left; right-- {
			subLen := right - left + 1
			if subLen <= resultLen {
				break
			}

			if char[left] == char[right] {
				ok := true
				i := left + 1
				j := right - 1
				for i <= j {
					if char[i] != char[j] {
						ok = false
						break
					} else {
						i++
						j--
					}
				}
				if ok && subLen > resultLen {
					resultLeft = left
					resultRight = right
					resultLen = subLen
				}
			}
		}
	}
	char = char[resultLeft : resultRight+1]

	return string(char)
}

func longestPalindromicSubstring2(s string) string {
	// TODO

	return ""
}
