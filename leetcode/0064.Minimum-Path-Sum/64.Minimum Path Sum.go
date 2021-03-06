package leetcode

func minPathSum(grid [][]int) int {
	// f(i, j) = min(f(i-1, j), f(i, j-1))
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 空间复杂度O(1)
func minPathSum2(grid [][]int) int {
	// f(i, j) = min(f(i-1, j), f(i, j-1))
	m := len(grid)
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func s(n int) {

}
