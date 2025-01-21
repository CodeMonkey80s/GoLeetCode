package solution1539

func findKthPositive(arr []int, k int) int {

	var m int
	seen := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		seen[arr[i]] = true
		m = max(arr[i], m)
	}

	var c int
	var number int
	var i int
	for {
		i++
		_, ok := seen[i]
		if !ok || i > m {
			number = i
			c++
		}
		if c == k {
			break
		}
	}
	return number
}
