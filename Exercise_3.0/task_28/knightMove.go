package task28

func KnightMove(n, m int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[1][1] = 1
	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			dp[i][j] = dp[i-2][j-1] + dp[i-1][j-2]
		}
	}
	return dp[n][m]
}
