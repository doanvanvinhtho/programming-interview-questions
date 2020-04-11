// Longest Palindromic Substring

// Given a string, find the longest substring which is palindrome.
// For example, if the given string is "forgeeksskeegfor", the output should be "geeksskeeg"

package bstring

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{
			in:  "forgeeksskeegfor",
			out: "geeksskeeg",
		},
		{
			in:  "forgeeks1skeegfor",
			out: "geeks1skeeg",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "1",
			out: "1",
		},
		{
			in:  "12",
			out: "1",
		},
		{
			in:  "121",
			out: "121",
		},
	}
	for i, c := range cases {
		r := longestPalindromicSubstring1(c.in)
		if c.out != r {
			t.Errorf("Wrong %v got %v", i, r)
		}
	}
}
