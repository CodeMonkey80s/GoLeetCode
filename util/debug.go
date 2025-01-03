package util

import (
	"cmp"
	"fmt"
	"slices"
)

func PrintOrderedMap[T1 cmp.Ordered, T2 any](m map[T1]T2) {
	keys := make([]T1, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, v := range keys {
		fmt.Printf("%v: %+v\n", v, m[v])
	}
}
