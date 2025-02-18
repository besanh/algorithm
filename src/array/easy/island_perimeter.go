package easy

/*
 * Problem: https://leetcode.com/problems/island-perimeter/?envType=problem-list-v2&envId=array
 * Solution: iterate each item, if item is 1, I check its neighbors, if neighbor is 0, I add 1 to perimeter
 * Time complexity: O(m*n), because I iterate rows and columns
 * Space complexity: O(1), because I iterate each item
 * DSA: normal
 */
func IslandPerimeter(grid [][]int) int {
	perimeter := 0
	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				// A square has 4 line segments
				perimeter += 4

				// Check top neighbor
				if r > 0 && grid[r-1][c] == 1 {
					perimeter -= 2
				}
				// Check left neighbor
				if c > 0 && grid[r][c-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}

	return perimeter
}
