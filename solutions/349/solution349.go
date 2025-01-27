package solution349

// ============================================================================
// 349. Intersection of Two Arrays
// URL: https://leetcode.com/problems/intersection-of-two-arrays/
// ============================================================================

/*

	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/349
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_intersectionV2
	Benchmark_intersectionV2-24    	 7600792	       154.8 ns/op	      56 B/op	       3 allocs/op
	Benchmark_intersectionV1
	Benchmark_intersectionV1-24    	 6605528	       172.7 ns/op	      56 B/op	       3 allocs/op
	PASS

*/

func intersectionV2(nums1 []int, nums2 []int) []int {

	freq := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		freq[nums1[i]] = 0
	}

	for i := 0; i < len(nums2); i++ {
		_, ok := freq[nums2[i]]
		if ok {
			freq[nums2[i]] = 1
		}
	}

	var output []int
	for k, c := range freq {
		if c == 1 {
			output = append(output, k)
		}
	}

	return output
}

func intersectionV1(nums1 []int, nums2 []int) []int {
	ans := []int{}
	m1 := make(map[int]int, len(nums1))
	m2 := make(map[int]int, len(nums2))
	for i := 0; i < len(nums1); i++ {
		val := nums1[i]
		_, ok := m1[val]
		if !ok {
			m1[val] = 1
		} else {
			m1[val]++
		}
	}
	for i := 0; i < len(nums2); i++ {
		val := nums2[i]
		_, ok := m2[val]
		if !ok {
			m2[val] = 1
		} else {
			m2[val]++
		}
	}
	if len(m1) < len(m2) {
		for k := range m1 {
			_, ok := m2[k]
			if !ok {
				continue
			}
			ans = append(ans, k)
		}
	} else {
		for k := range m2 {
			_, ok := m1[k]
			if !ok {
				continue
			}
			ans = append(ans, k)
		}
	}
	return ans
}
