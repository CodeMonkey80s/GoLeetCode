package solution3354

/*
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/3354
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_countValidSelectionsV1
	Benchmark_countValidSelectionsV1-24    	22890769	        48.45 ns/op	      96 B/op	       2 allocs/op
	Benchmark_countValidSelectionsV2-24    	 9561316	       119.7 ns/op	     192 B/op	       4 allocs/op
	Benchmark_countValidSelectionsV3-24    	78240026	        15.28 ns/op	       0 B/op	       0 allocs/op
	PASS
*/

func countValidSelectionsV1(nums []int) int {

	prefixSumL := make([]int, len(nums))
	prefixSumR := make([]int, len(nums))

	var total int
	prefixSumL[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSumL[i] = prefixSumL[i-1] + nums[i]
		total = max(total, prefixSumL[i])
	}

	prefixSumR[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		prefixSumR[i] = prefixSumR[i+1] + nums[i]
	}

	var valid int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			switch {
			case prefixSumL[i]-prefixSumR[i] == 0:
				valid += 2
			case prefixSumL[i]+1 == prefixSumR[i]:
				valid += 1
			case prefixSumL[i] == prefixSumR[i]+1:
				valid += 1
			}
		}
	}

	return valid
}

func countValidSelectionsV3(nums []int) int {

	sum := func(idx int) int {
		var sL int
		var sR int
		for i := idx; i < len(nums); i++ {
			sR += nums[i]
		}
		for i := idx; i >= 0; i-- {
			sL += nums[i]
		}

		if sL == sR {
			return 2
		}

		if sL-1 == sR || sR-1 == sL {
			return 1
		}

		return 0
	}

	var valid int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			valid += sum(i)
		}
	}

	return valid
}

func countValidSelectionsV2(nums []int) int {

	ping := func(curr int, d int) int {

		numsCurrent := make([]int, len(nums))
		copy(numsCurrent, nums)

		for {
			if curr < 0 || curr >= len(nums) {
				break
			}

			switch {
			case numsCurrent[curr] == 0:
				curr += d
			case numsCurrent[curr] > 0:
				numsCurrent[curr]--
				d = -d
				curr += d
			}
		}

		for i := 0; i < len(numsCurrent); i++ {
			if numsCurrent[i] != 0 {
				return 0
			}
		}

		return 1
	}

	var valid int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			valid += ping(i, -1)
			valid += ping(i, 1)
		}
	}

	return valid
}
