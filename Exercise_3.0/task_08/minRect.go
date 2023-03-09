package task08

func MinRect(sqCnt int, coords [][2]int) [4]int {
	x1, y1, x2, y2 := coords[0][0], coords[0][1], coords[0][0], coords[0][1]
	for _, c := range coords {
		if x1 > c[0] {
			x1 = c[0]
		}
		if y1 > c[1] {
			y1 = c[1]
		}
		if x2 < c[0] {
			x2 = c[0]
		}
		if y2 < c[1] {
			y2 = c[1]
		}
	}
	return [4]int{x1, y1, x2, y2}
}
