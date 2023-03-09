package task26

func CheapestWay(n, m int, ssi [][]int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = ssi[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + ssi[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + ssi[i][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j]+ssi[i][j], dp[i][j-1]+ssi[i][j])
		}
	}
	return dp[n-1][m-1]
}

func min(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if m > v {
			m = v
		}
	}
	return m
}
