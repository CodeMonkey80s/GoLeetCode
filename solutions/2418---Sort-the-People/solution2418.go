package solution2418

import (
	"sort"
)

// ============================================================================
// 2418. Sort the People
// URL: https://leetcode.com/problems/sort-the-people/
// ============================================================================

/*
	$ go test -bench=. -benchmem
	goos: linux
	goarch: amd64
	pkg: GoLeetCode/solutions/2418---Sort-the-People
	cpu: 13th Gen Intel(R) Core(TM) i7-13700K
	Benchmark_sortPeople-24    	 2977778	       397.2 ns/op	     416 B/op	       5 allocs/op
	PASS
*/

func sortPeople(names []string, heights []int) []string {
	type Person struct {
		Name   string
		Height int
	}
	people := make([]Person, 0, len(names))
	for i := range names {
		p := Person{
			Name:   names[i],
			Height: heights[i],
		}
		people = append(people, p)
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Height >= people[j].Height
	})
	sorted := make([]string, 0, len(names))
	for _, v := range people {
		sorted = append(sorted, v.Name)
	}
	return sorted
}
