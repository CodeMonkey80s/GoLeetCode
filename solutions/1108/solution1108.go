package solution1108

import "strings"

// ============================================================================
// 1108. Defanging an IP Address
// URL: https://leetcode.com/problems/defanging-an-ip-address/
// ============================================================================

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
