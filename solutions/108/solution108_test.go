package solution108

import (
	"strconv"
)

var testCases = []struct {
	Input  []int
	Output []string
}{
	// Mandatory Test Cases
	{
		Input:  []int{-10, -3, 0, 5, 9},
		Output: []string{"0", "-3", "9", "-10", "null", "5"},
	},
	// Additional my custom cases
}

func convert(tree *TreeNode) []string {
	s := []string{}
	sv := ""
	if tree == nil {
		s = append([]string{}, "null")
		return s
	}
	s1 := convert(tree.Left)
	s = append(s, s1...)
	sv = strconv.Itoa(tree.Val)
	s = append(s, sv)
	s2 := convert(tree.Right)
	s = append(s, s2...)
	sv = strconv.Itoa(tree.Val)
	s = append(s, sv)
	return s
}

/*
func Test_sorted(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			head := sortedArrayToBST(tc.Input)
			output := convert(head)
			if !reflect.DeepEqual(tc.Output, output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortedArrayToBST(testCases[0].Input)
	}
}
*/
