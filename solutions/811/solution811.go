package solution811

import (
	"strconv"
	"strings"
)

// ============================================================================
// 811. Subdomain Visit Count
// URL: https://leetcode.com/problems/subdomain-visit-count/description/
// ============================================================================

func subdomainVisits(cpdomains []string) []string {

	visits := make(map[string]int)
	for _, info := range cpdomains {
		parts := strings.Fields(info)
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]
		subs := strings.Split(domain, ".")
		for i := 0; i < len(subs); i++ {
			subdomain := strings.Join(subs[i:], ".")
			visits[subdomain] += count
		}
	}

	var subdomains []string
	for domain, count := range visits {
		s := strconv.Itoa(count) + " " + domain
		subdomains = append(subdomains, s)
	}

	return subdomains
}
