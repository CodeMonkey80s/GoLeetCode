package solution860

func lemonadeChange(bills []int) bool {

	collected := make(map[int]int)
	for i := 0; i < len(bills); i++ {

		value := bills[i]
		collected[value]++

		if value == 5 {
			continue
		}

		d := value - 5
		for d > 0 {
			if d >= 10 && collected[10] > 0 {
				collected[10]--
				d -= 10
			} else if collected[5] > 0 {
				collected[5]--
				d -= 5
			} else {
				return false
			}
		}
	}

	return true
}
