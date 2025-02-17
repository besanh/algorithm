package medium

/*
 * Problem: https://leetcode.com/problems/number-of-islands/description/?envType=problem-list-v2&envId=array
 * Time complexity: O(m*n), because I visit each element once
 * Space complexity: O(m*n), because I use a 2D array to mark visited elements
 * DSA: DFS(recursive or stack), BFS(recursive or stack)
 * My solution: use DFS combine with recursive, easy to implement
 * Explaination: I iterate each item in 2D array and check if it is '1', if yes, I mark it as '0' and
 * call the DFS function to visit its neighbors, if a neighbor is also '1', I call the DFS function again
 * until all neighbors are visited
 */
func NumIslands(grid [][]byte) int {
	// Write your code here
	totalIslands := 0
	rows, cols := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// I check grid[row][col] == '0' because I have already mark it as '0' in line 24
		if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == '0' {
			return
		}

		// Mark as visited
		grid[row][col] = '0'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Start with condition '1', meaning I met a land
			if grid[r][c] == '1' {
				totalIslands++
				dfs(r, c)
			}
		}
	}

	return totalIslands
}
