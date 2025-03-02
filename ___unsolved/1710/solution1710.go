package solution1710

import (
	"fmt"
	"slices"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {

	slices.SortFunc(boxTypes, func(a, b []int) int {
		return b[1] - a[1]
	})

	fmt.Println(boxTypes)

	var units int
	for i := 0; i < len(boxTypes); i++ {

		if boxTypes[i][0] < truckSize {
			c := boxTypes[i][0] * boxTypes[i][1]
			units += c
			truckSize -= c
		}
	}

	return units
}
