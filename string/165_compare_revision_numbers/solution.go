package string

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")

	maxL := max(len(v1s), len(v2s))

	for i := range maxL {
		a, b := 0, 0

		if i < len(v1s) {
			a, _ = strconv.Atoi(v1s[i])
		}

		if i < len(v2s) {
			b, _ = strconv.Atoi(v2s[i])
		}

		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}

	return 0
}
