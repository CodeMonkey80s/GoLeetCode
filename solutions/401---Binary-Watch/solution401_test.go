package solution401

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []string
}{
	{
		Input:  1,
		Output: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
	},
	{
		Input:  2,
		Output: []string{"0:03", "0:05", "0:06", "0:09", "0:10", "0:12", "0:17", "0:18", "0:20", "0:24", "0:33", "0:34", "0:36", "0:40", "0:48", "1:01", "1:02", "1:04", "1:08", "1:16", "1:32", "2:01", "2:02", "2:04", "2:08", "2:16", "2:32", "3:00", "4:01", "4:02", "4:04", "4:08", "4:16", "4:32", "5:00", "6:00", "8:01", "8:02", "8:04", "8:08", "8:16", "8:32", "9:00", "10:00"},
	},
	{
		Input:  0,
		Output: []string{"0:00"},
	},
	{
		Input:  9,
		Output: []string{},
	},
}

func Test_readBinaryWatch(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := readBinaryWatch(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_readBinaryWatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = readBinaryWatch(testCases[0].Input)
	}
}
