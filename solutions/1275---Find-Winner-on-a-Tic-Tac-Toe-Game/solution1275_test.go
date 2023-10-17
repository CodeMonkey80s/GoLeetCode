package solution1275

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output string
}{
	{
		Input: [][]int{
			{0, 0},
			{2, 0},
			{1, 1},
			{2, 1},
			{2, 2},
		},
		Output: "A",
	},
	{
		Input: [][]int{
			{0, 0},
			{1, 1},
			{0, 1},
			{0, 2},
			{1, 0},
			{2, 0},
		},
		Output: "B",
	},
	{
		Input: [][]int{
			{0, 0},
			{1, 1},
			{2, 0},
			{1, 0},
			{1, 2},
			{2, 1},
			{0, 1},
			{0, 2},
			{2, 2},
		},
		Output: "Draw",
	},
	{
		Input: [][]int{
			{0, 0},
			{1, 1},
		},
		Output: "Pending",
	},
	{
		Input: [][]int{
			{1, 2},
			{2, 1},
			{1, 0},
			{0, 0},
			{0, 1},
			{2, 0},
			{1, 1},
		},
		Output: "A",
	},
}

func Test_tictactoe(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := tictactoe(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_tictactoe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tictactoe(testCases[0].Input)
	}
}
