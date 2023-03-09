package task02

func PrettyString(s string, swap int) (mpr int) {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	for i := 97; i <= 122; i++ {
		curSwap := swap
		curMax := 0
		for l, r := 0, 0; r < len(s) && l < len(s); {
			if int(s[r]) == i {
				curMax++
				r++
			} else if int(s[r]) != i && curSwap != 0 {
				curMax++
				r++
				curSwap--
			} else if int(s[r]) != i && curSwap == 0 {
				if mpr < curMax {
					mpr = curMax
				}
				if int(s[l]) == i {
					l++
					curMax--
				} else {
					l++
					curSwap++
					curMax--
				}
			}
		}
		if curMax > mpr {
			mpr = curMax
		}
	}
	return
}
