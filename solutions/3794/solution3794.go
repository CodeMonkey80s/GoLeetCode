package solution3794

import (
	"slices"
	"strings"
)

func reversePrefix(s string, k int) string {
	b := []byte(s[:k])
	slices.Reverse(b)

	var sb strings.Builder
	sb.Write(b)
	sb.Write([]byte(s[k:]))

	return sb.String()
}
