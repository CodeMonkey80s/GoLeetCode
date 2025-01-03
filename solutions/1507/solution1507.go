package solution1507

import (
	"fmt"
	"strings"
)

// ============================================================================
// 1507. Reformat Date
// URL: https://leetcode.com/problems/reformat-date/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_reformatDate-24    	 2624359	       477.9 ns/op	     714 B/op	       7 allocs/op
	PASS

*/

func reformatDate(date string) string {
	sb := strings.Builder{}
	month := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	parts := strings.Fields(date)
	sb.WriteString(parts[2])
	sb.WriteString("-")
	v, ok := month[parts[1]]
	if !ok {
		panic("value not found!")
	}
	sb.WriteString(v)
	sb.WriteString("-")
	num := parts[0]
	switch {
	case 'a' <= num[1] && num[1] <= 'z':
		num = num[:1]
	case 'a' <= num[2] && num[2] <= 'z':
		num = num[:2]
	}
	num = fmt.Sprintf("%02s", num)
	sb.WriteString(num)
	return sb.String()
}
