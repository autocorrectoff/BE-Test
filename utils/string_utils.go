package utils

import "regexp"

// LongestString -> returns the longest string in a slice of strings
func LongestString(s []string) string {
	longest, length := "", 0
	for _, word := range s {
		if len(word) > length {
			longest, length = word, len(word)
		}
	}
	return longest
}

// ContainsNumerics -> returns true if string contains one or more numerics
func ContainsNumerics(str string) bool {
	return regexp.MustCompile(`\d`).MatchString(str)
}
