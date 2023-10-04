package solution350

// ============================================================================
// 350. Intersection of Two Arrays II
// URL: https://leetcode.com/problems/intersection-of-two-arrays-ii/
// ============================================================================

/*

	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_intersection-24    	 7460503	       171.3 ns/op	      72 B/op	       4 allocs/op
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
		for k, v1 := range m1 {
			v2, ok := m2[k]
			minval := v2
			if !ok {
				continue
			}
			if v1 < minval {
				minval = v1
			}
			for i := 0; i < minval; i++ {
				ans = append(ans, k)
			}
		}
	} else {
		for k, v2 := range m2 {
			v1, ok := m1[k]
			minval := v1
			if !ok {
				continue
			}
			if v2 < minval {
				minval = v2
			}
			for i := 0; i < minval; i++ {
				ans = append(ans, k)
			}
		}
	}
	return ans
}
