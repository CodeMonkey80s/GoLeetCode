package solution1694

import (
	"strings"
)

func reformatNumber(number string) string {

	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")

	left := len(number) % 3

	var sb strings.Builder
	for i := 0; i < len(number)-left; i += 3 {
		sb.WriteString(number[i : i+3])
		if i+3 < len(number) {
			sb.WriteByte('-')
		}
	}

	return sb.String()
}
