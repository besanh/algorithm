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
	for i := 0; i < m; i++ {
		for j := 0; j < n-1; j++ {
			if grid[j][i] > grid[j+1][i] {
				return "NO"
			}
		}
	}
	return "YES"
}
