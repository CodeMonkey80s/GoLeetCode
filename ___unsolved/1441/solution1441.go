package solution1441

import (
	"reflect"
)

func buildArray(target []int, n int) []string {

	var output []string
	var stack []int
	var i int
	for i <= n {

		i++
		stack = append(stack, i)
		output = append(output, "Push")
		if reflect.DeepEqual(stack, target) {
			break
		}

		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
			output = append(output, "Pop")
		}

		if reflect.DeepEqual(stack, target) {
			break
		}
	}

	return output
}
