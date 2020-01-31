// Reverse words in a given string

// Example: Let the input string be “i like this program very much”.
// The function should change the string to “much very program this like i”

package bstring

import "testing"

func TestReverseWordsInString(t *testing.T) {
	s1 := reverseWordsInString1("i like this program very much")
	if s1 != "much very program this like i" {
		t.Errorf("Wrong 1, [%v]", s1)
	}

	s2 := reverseWordsInString2("i like this program very much")
	if s2 != "much very program this like i" {
		t.Errorf("Wrong 2, [%v]", s2)
	}
}

// goos: darwin
// goarch: amd64
// BenchmarkReverseWordsInString1-4   	  521012	      1943 ns/op	     672 B/op	      47 allocs/op
// BenchmarkReverseWordsInString2-4   	 2080260	       545 ns/op	      32 B/op	       1 allocs/op

func BenchmarkReverseWordsInString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsInString1("i like this program very much")
	}
}

func BenchmarkReverseWordsInString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsInString2("i like this program very much")
	}
}
