package task29

import "fmt"

func Cafe(cnt int, si []int) (amount int, coup [2]int, day int) {
	dp := make([]int, cnt)
	dp[0] = si[0]
	for i := 1; i < cnt; i++ {
		dp[i] = dp[i-1] + si[i]
	}
	for i := 0; i < cnt; i++ {
		if si[i] > 100 {
			coup[0]++
		}
		if coup[0] > 0 {
			coup[1]++
			coup[0]--
		}
	}
	fmt.Println(dp, coup)
	return
}
