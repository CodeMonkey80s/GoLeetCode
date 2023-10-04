package solution190

// ============================================================================
// 190. Reverse Bits
// https://leetcode.com/problems/reverse-bits/
// ============================================================================

func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32; i++ {
		rev = rev << 1
		if (num & 1) == 1 {
			rev ^= 1
		}
		num = num >> 1
	}
	return rev
}

/*
func reverseBits(num uint32) uint32 {
	var n uint32
	var a uint32 = 32
	var b uint32 = 0
	var v uint32
	for a > 0 {
		v = num & (1 << (a - 1))
		if v > 0 {
			n += uint32(math.Pow(2, float64(b)))
		}
		a--
		b++
	}

	return n
}
*/
