package task20

func PyramidSort(cnt int, si []int) []int {
	for i := len(si)/2 - 1; i >= 0; i-- {
		si = heapSort(si, i, len(si))
	}

	for i := len(si) - 1; i >= 1; i-- {
		si[0], si[i] = si[i], si[0]
		si = heapSort(si, 0, i)
	}
	return si
}

func heapSort(si []int, i int, l int) []int {
	max := 0
	for i*2+1 < l {
		if i*2+1 == l-1 || si[i*2+1] > si[i*2+2] {
			max = i*2 + 1
		} else {
			max = i*2 + 2
		}

		if si[i] < si[max] {
			si[i], si[max] = si[max], si[i]
			i = max
		} else {
			break
		}
	}
	return si
}
