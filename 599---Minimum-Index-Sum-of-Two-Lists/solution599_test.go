package solution599

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB []string
	Output []string
}{
	// Mandatory Test Cases
	{
		InputA: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		InputB: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		Output: []string{"Shogun"},
	},
	{
		InputA: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		InputB: []string{"KFC", "Shogun", "Burger King"},
		Output: []string{"Shogun"},
	},
	{
		InputA: []string{"happy", "sad", "good"},
		InputB: []string{"sad", "happy", "good"},
		Output: []string{"happy", "sad"},
	},
	// Additional my custom cases
	{
		InputA: []string{"Shogun", "Piatti", "Tapioca Express", "Burger King", "KFC"},
		InputB: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		Output: []string{"Piatti"},
	},
}

func Test_findRestaurant(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findRestaurant_first_attempt(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findRestaurant_first(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findRestaurant_first_attempt(testCases[0].InputA, testCases[0].InputB)
	}
}
