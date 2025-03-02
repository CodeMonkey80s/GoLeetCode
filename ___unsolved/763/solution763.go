package solution763

import (
	"fmt"
)

func partitionLabels(s string) []int {

	var output []int
	var b int

loop:
	for len(s) > 0 {

		fmt.Printf("chk: %q, %d\n", s, len(s))

		b = len(s)
		for {

			b--
			if b <= 0 {
				break loop
			}

			// if b == 0 {
			// 	fmt.Printf("s1: %q, %d\n", s[:b], b)
			// 	output = append(output, len(s))
			// 	break loop
			// }

			if s[0] == s[b] {
				fmt.Printf("cha: %c %c\n", s[0], s[b])
				e := b + 1
				fmt.Printf("cut: %q, %d\n", s[:e], e)
				output = append(output, len(s[:e]))
				s = s[e:]
				fmt.Printf("nxt: %q, len: %d\n\n", s, len(s))
				goto loop
			}

		}
	}

	return output
}
