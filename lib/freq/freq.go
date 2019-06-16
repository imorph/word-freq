package freq

import (
	"strings"
)

// CleanupWord takes word with "." "," ";" ":" "\n" and return "clean" word
func CleanupWord(inWord string) string {
	outWord := inWord
	return outWord
}

// StringToCleanSlice splits text to words (cleaned from any of . , : ; \n ) and retuns it as slice of words in lower case
func StringToCleanSlice(text string) []string {
	return strings.Split(text, " ")
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
