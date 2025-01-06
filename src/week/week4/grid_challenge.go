package week4

import (
	"sort"
	"strings"
)

/*
* Sort the grid in lexicographical order and return it
* Begin from the row
* Then from the column, if row[i][j] > row[i+1][j](meaning it don't sort) => return NO
 */
func GridChallenge(grid []string) string {
	for i, char := range grid {
		row := strings.Split(char, "")
		sort.Strings(row)
		grid[i] = strings.Join(row, "")
	}

	n, m := len(grid), len(grid[0])
	// Don't need to check the last row
	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			// Check in same column, if next row is smaller, it means not sorted
			if grid[i][j] > grid[i+1][j] {
				return "NO"
			}
		}
	}
	return "YES"
}
