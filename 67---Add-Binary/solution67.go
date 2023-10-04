package solution67

import (
	"fmt"
	"strconv"
)

// ============================================================================
// 67. Add Binary
// URL: https://leetcode.com/problems/add-binary/
// ============================================================================

func addBinary_stdlib(a string, b string) string {
	va, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		panic(err)
	}
	vb, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	sum := va + vb
	s := fmt.Sprintf("%b", sum)
	return s
}

func addBinary_my(a string, b string) string {
	sa := len(a)
	sb := len(b)
	m := sa
	if sb > sa {
		m = sb
	}
	s := make([]byte, m+1)
	var va int
	var vb int
	var carry int
	var i int = 1
	var sum int
	for {
		va = 0
		vb = 0
		if i <= sa {
			va = int(a[sa-i]) - 48
		}
		if i <= sb {
			vb = int(b[sb-i]) - 48
		}
		sum = va + vb + carry
		switch {
		case sum == 0:
			s[m] = '0'
			carry = 0
		case sum == 1:
			s[m] = '1'
			carry = 0
		case sum == 2:
			s[m] = '0'
			carry = 1
		case sum == 3:
			s[m] = '1'
			carry = 1
		}
		i++
		if i > sa+1 && i > sb+1 {
			break
		}
		m--
	}
	if s[0] == '0' {
		s = s[1:]
	}
	return string(s)
}
