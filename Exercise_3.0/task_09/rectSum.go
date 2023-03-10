package task09

func RectSum(base [3]int, ssi [][]int, sReq [][4]int) []int {
	n, m, r := base[0], base[1], base[2]
	pref := make([][]int, n)
	for i := range pref {
		pref[i] = make([]int, m)
	}

	pref[0][0] = ssi[0][0]
	for i := 1; i < n; i++ {
		pref[i][0] = pref[i-1][0] + ssi[i][0]
	}

	for i := 1; i < m; i++ {
		pref[0][i] = pref[0][i-1] + ssi[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			pref[i][j] = pref[i-1][j] + pref[i][j-1] - pref[i-1][j-1] + ssi[i][j]
		}
	}

	rs := make([]int, r)
	for i, v := range sReq {
		f := pref[v[2]-1][v[3]-1]

		l := 0
		if v[1]-1 > 0 {
			l = pref[v[2]-1][v[1]-2]
		}

		t := 0
		if v[0]-1 > 0 {
			t = pref[v[0]-2][v[3]-1]
		}

		p := 0
		if v[0]-1 > 0 && v[1]-1 > 0 {
			p = pref[v[0]-2][v[1]-2]
		}

		rs[i] = f - l - t + p
	}
	return rs
}
