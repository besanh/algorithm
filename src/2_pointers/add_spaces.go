package pointers

import (
	"sort"
	"strings"
)

// notme
func AddSpaces(s string, spaces []int) string {
	var sb strings.Builder
	spaceIdx := 0
	n := len(spaces)

	sort.Ints(spaces)

	for i := 0; i < len(s); i++ {
		if spaceIdx < n && i == spaces[spaceIdx] {
			sb.WriteByte(' ')
		}
		sb.WriteByte(s[i])
	}

	if spaceIdx < n && spaces[spaceIdx] == len(s) {
		sb.WriteByte(' ')
	}

	return sb.String()
}
