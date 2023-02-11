package utils

// ReplaceAtIndex -> replaces string at index
func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
