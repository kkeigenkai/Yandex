package task05

func GoodString(cnt int, si []int) int {
	good := 0
	min := si[0]
	for i := 1; i < len(si); i++ {
		if min > si[i] {
			min = si[i]
		}
		good += min
		min = si[i]
	}
	return good
}
