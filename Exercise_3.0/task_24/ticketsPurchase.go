package task24

func TicketPurchase(cnt int, que [][3]int) int {
	if cnt == 1 {
		return min(que[0][0], que[0][1], que[0][2])
	}
	que = append([][3]int{{0, 0, 0}}, que...)
	t := make([]int, cnt+1)
	t[0] = 0
	t[1] = que[1][0]
	t[2] = min(que[1][0]+que[2][0], que[1][1])

	for i := 3; i <= cnt; i++ {
		t[i] = min(t[i-1]+que[i][0], t[i-2]+que[i-1][1], t[i-3]+que[i-2][2])
	}
	return t[cnt]
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
