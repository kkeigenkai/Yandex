package task23

func Calc(n int) (int, []int) {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		m := dp[i-1] + 1
		if i%2 == 0 {
			m = min(m, dp[i/2]+1)
		}
		if i%3 == 0 {
			m = min(m, dp[i/3]+1)
		}
		dp[i] = m
	}

	si := make([]int, dp[n]+1)
	si[len(si)-1] = n
	for i, j := n, dp[n]-1; i > 1; j-- {
		if dp[i] == dp[i-1]+1 {
			i -= 1
		} else if i%2 == 0 && dp[i] == dp[i/2]+1 {
			i /= 2
		} else {
			i /= 3
		}
		si[j] = i
	}
	return dp[n], si
}

func min(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}
