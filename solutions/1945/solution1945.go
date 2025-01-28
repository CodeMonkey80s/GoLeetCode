package solution1945

import (
	"fmt"
	"strconv"
	"strings"
)

func getLucky(s string, k int) int {

	b := []byte(s)
	for i := 0; i < len(s); i++ {
		b[i] = b[i] - 'a' + 1
	}

	var sum int
	var sb strings.Builder
	for i := 0; i < k; i++ {
		sb.Reset()
		for j := 0; j < len(b); j++ {
			v := int(b[j] - '0')
			sb.WriteString(strconv.Itoa(v))
			sum += v

		}

		b = []byte(sb.String())
	}

	fmt.Println(sum)

	return sum
}
