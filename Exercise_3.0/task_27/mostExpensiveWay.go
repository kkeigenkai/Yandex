package task27

import (
	"strings"
)

func MostExpensiveWay(n, m int, ssi [][]int) (int, string) {
	dp := make([][]int, n)
	for i := range dp {
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
			dp[i][j] = max(dp[i-1][j]+ssi[i][j], dp[i][j-1]+ssi[i][j])
		}
	}
	prev := make([]string, 0)
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if j > 0 && dp[i][j] == ssi[i][j]+dp[i][j-1] {
			prev = append(prev, "R")
			j--
		} else if i > 0 {
			prev = append(prev, "D")
			i--
		} else {
			break
		}
	}
	for i, j := 0, len(prev)-1; i < len(prev)/2; i, j = i+1, j-1 {
		prev[i], prev[j] = prev[j], prev[i]
	}
	return dp[n-1][m-1], strings.Join(prev, " ")
}

func max(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if m < v {
			m = v
		}
	}
	return m
}
