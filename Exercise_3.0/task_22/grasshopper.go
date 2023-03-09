package task22

const MAXSIZE = 31

func Grasshopper(size, jump int) int {
	if size <= 2 {
		return 1
	}
	dp := make([]int, MAXSIZE)
	dp[1], dp[2] = 1, 1
	i := 3
	for ; i <= jump; i++ {
		dp[i] = 2 * dp[i-1]
	}

	for ; i <= size; i++ {
		dp[i] = 2*dp[i-1] - dp[i-jump-1]
	}
	return dp[size]
}
