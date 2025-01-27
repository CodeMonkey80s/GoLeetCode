package solution929

import (
	"strings"
)

func numUniqueEmails(emails []string) int {

	unique := make(map[string]struct{})
	for _, email := range emails {

		parts := strings.Split(email, "@")
		name := strings.ReplaceAll(parts[0], ".", "")
		idx := strings.IndexByte(name, '+')
		if idx != -1 {
			name = name[:idx]
		}

		e := name + "@" + parts[1]
		unique[e] = struct{}{}
	}

	return len(unique)
}
