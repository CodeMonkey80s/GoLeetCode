package solution1413

func minStartValue(nums []int) int {

	var s int
	var m int
	for _, num := range nums {
		s += num
		m = min(m, s)
	}
	return 1 - m

}
