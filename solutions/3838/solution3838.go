package solution3838

import "strings"

func mapWordWeights(words []string, weights []int) string {
	var sb strings.Builder

	for i := 0; i < len(words); i++ {
		var weight int
		for j := 0; j < len(words[i]); j++ {
			weight += weights[int(words[i][j]-'a')]
		}
		sb.WriteByte(byte('z' - weight%26))
	}

	return sb.String()
}
