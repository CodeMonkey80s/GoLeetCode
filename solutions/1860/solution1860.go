package solution1860

func memLeak(memory1 int, memory2 int) []int {

	t := 1
loop:
	for {
		switch {
		case memory1 >= memory2:
			if memory1-t < 0 {
				break loop
			}
			memory1 -= t
		default:
			if memory2-t < 0 {
				break loop
			}
			memory2 -= t
		}
		t++
	}

	return []int{t, memory1, memory2}
}
