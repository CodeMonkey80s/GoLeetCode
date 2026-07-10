package solution3950

import "strconv"

func consecutiveSetBits(n int) bool {
	var flag bool
	bits := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(bits)-1; i++ {
		if bits[i] == '1' && bits[i+1] == '1' {
			if flag == false {
				flag = true
			} else {
				flag = false
				break
			}
		}
	}

	return flag
}
