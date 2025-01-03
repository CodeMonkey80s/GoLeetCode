package solution349

// ============================================================================
// 349. Intersection of Two Arrays
// URL: https://leetcode.com/problems/intersection-of-two-arrays/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_intersection-24    	 9420332	       134.6 ns/op	      56 B/op	       3 allocs/op
	PASS

*/

func intersection(nums1 []int, nums2 []int) []int {
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
