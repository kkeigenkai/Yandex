package prettystring

import "fmt"

func PrettyString(s string, swap int) (mpr int) {
	mi := make(map[rune][]int)

	for i, c := range s {
		mi[c] = append(mi[c], i)
	}

	mm := make(map[rune]int)

	for k, v := range mi {
		curM := 0
		for i := len(v) - 1; i > 0; i-- {
			curM += v[i] - v[i-1] + 1
		}
		mm[k] = curM
	}
	fmt.Println(mi)
	fmt.Println(mm)
	return
}

func minMax(m map[rune]int) (min, max int) {
	min = 1
	max = 0
	for _, v := range m {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return
}
