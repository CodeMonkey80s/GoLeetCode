package solution28

import "strings"

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if s := haystack[i : i+len(needle)]; s == needle {
			return i
		}
	}
	return -1
}

func strStr_stdlib(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
