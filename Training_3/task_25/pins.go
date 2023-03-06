package task25

import (
	"sort"
)

func Pins(cnt int, pins []int) int {
	np := append([]int{0}, pins...)
	sort.Slice(np, func(i, j int) bool { return np[i] < np[j] })
	if len(np) == 3 {
		return np[2] - np[1]
	}
	dp := make([]int, cnt+1)
	dp[2], dp[3] = np[2]-np[1], np[3]-np[1]
	for i := 4; i <= cnt; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + np[i] - np[i-1]
	}
	return dp[cnt]
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
