package solution1805

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  "a123bc34d8ef34",
		Output: 3,
	},
	{
		Input:  "leet1234code234",
		Output: 2,
	},
	// Additional my custom cases
	{
		Input:  "a1b01c001",
		Output: 1,
	},
	{
		Input:  "2393706880236110407059624696967828762752651982730115221690437821508229419410771541532394006597463715513741725852432559057224478815116557380260390432211227579663571046845842281704281749571110076974264971989893607137140456254346955633455446057823738757323149856858154529105301197388177242583658641529908583934918768953462557716z97438020429952944646288084173334701047574188936201324845149110176716130267041674438237608038734431519439828191344238609567530399189316846359766256507371240530620697102864238792350289978450509162697068948604722646739174590530336510475061521094503850598453536706982695212493902968251702853203929616930291257062173c79487281900662343830648295410",
		Output: 3,
	},
}

func Test_numDifferent(t *testing.T) {
	var label string
	for _, tc := range testCases {
		name := tc.Input
		if len(tc.Input) > 20 {
			name = tc.Input[0:20] + "..."
		}
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numDifferentIntegers(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numDifferent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numDifferentIntegers(testCases[0].Input)
	}
}
