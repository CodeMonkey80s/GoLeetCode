package solution3412

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int64
}{
	{
		Input:  "aczzx",
		Output: 5,
	},
	{
		Input:  "abcdef",
		Output: 0,
	},
	{
		Input:  "eockppxdqclkhjgvnw",
		Output: 50,
	},
	{
		Input:  "azapfwonwwcdagew",
		Output: 3,
	},
}

func Test_calculateScore(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := calculateScore(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_mirror(t *testing.T) {
	testCases := []struct {
		Input  byte
		Output byte
	}{
		{
			Input:  'a',
			Output: 'z',
		},
		{
			Input:  'm',
			Output: 'n',
		},
		{
			Input:  'z',
			Output: 'a',
		},
		{
			Input:  'c',
			Output: 'x',
		},
		{
			Input:  'x',
			Output: 'c',
		},
		{
			Input:  'y',
			Output: 'b',
		},
		{
			Input:  'b',
			Output: 'y',
		},
		{
			Input:  'k',
			Output: 'p',
		},
		{
			Input:  'p',
			Output: 'k',
		},
	}
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := mirror(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
