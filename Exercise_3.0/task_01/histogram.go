package task01

import (
	"sort"
)

func Histogram(s string) string {
	m := make(map[rune]int)
	for _, c := range s {
		if c != ' ' && c != '\n' {
			m[c]++
		}
	}

	sk := make([]rune, 0, len(m))
	for k := range m {
		sk = append(sk, k)
	}
	sort.Slice(sk, func(i, j int) bool { return sk[i] < sk[j] })

	res := make([]rune, 0)
	_, max := minMax(m)
	for ; max != 0; max-- {
		for _, v := range sk {
			if m[v] >= max {
				res = append(res, '#')
			} else {
				res = append(res, ' ')
			}
		}
		res = append(res, '\n')
	}
	return string(append(res, sk...))
}

func minMax(m map[rune]int) (min, max int) {
	min = 1
	max = 0
	for _, v := range m {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}
