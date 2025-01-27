package solution929

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
		Output: 2,
	},
	{
		Input:  []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"},
		Output: 3,
	},
	{
		Input:  []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"},
		Output: 2,
	},
}

func Test_numUniqueEmails(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numUniqueEmails(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
