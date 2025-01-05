package solution1310

// ============================================================================
// 1310. XOR Queries of a Subarray
// URL: https://leetcode.com/problems/xor-queries-of-a-subarray/d
// ============================================================================

func xorQueries(arr []int, queries [][]int) []int {

	prefix := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

	output := make([]int, len(queries))
	for i, q := range queries {
		output[i] = prefix[q[1]+1] ^ prefix[q[0]]
	}

	return output
}
