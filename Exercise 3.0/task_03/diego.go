package collectordiego

import (
	"sort"
)

func Diego(diegoCnt int, diegoS []int, collCnt int, collS []int) []int {
	if diegoCnt == 0 || collCnt == 0 {
		return []int{0}
	}

	sort.Slice(diegoS, func(i, j int) bool { return diegoS[i] < diegoS[j] })

	cleanS := []int{diegoS[0]}
	for i := 1; i < len(diegoS); i++ {
		if diegoS[i] != diegoS[i-1] {
			cleanS = append(cleanS, diegoS[i])
		}
	}

	res := make([]int, 0, collCnt)
	for _, v := range collS {
		if v <= cleanS[0] {
			res = append(res, 0)
			continue
		}
		l, r, mid, last := 0, len(cleanS)-1, 0, 0
		for l <= r {
			mid = (l + r) / 2
			if cleanS[mid] > v {
				r = mid - 1
			} else if cleanS[mid] < v {
				if cleanS[mid] > cleanS[last] {
					last = mid
				}
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		res = append(res, last+1)
	}
	return res
}
