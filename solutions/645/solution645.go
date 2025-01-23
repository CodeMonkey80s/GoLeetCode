package solution645

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/645
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_findErrorNums
	Benchmark_findErrorNums-24    	15787155	        67.22 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func findErrorNums(nums []int) []int {

	var a int
	var b int

	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			a = num
		}
	}

	for i := 1; i <= len(nums); i++ {
		_, ok := freq[i]
		if !ok {
			b = i
			break
		}
	}

	return []int{a, b}
}
