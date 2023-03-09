package task04

func ControlWork(sc, vc, pRow, pSide int) (vRow, vSide int) {
	pPos := pRow*2 - pSide%2
	vPosL := pPos - vc
	vPosR := pPos + vc
	vRowL := vPosL/2 + vPosL%2
	vRowR := vPosR/2 + vPosR%2
	if vPosL < 1 && vPosR > sc {
		return -1, 0
	}
	psl := 1
	if vPosL%2 == 0 {
		psl = 2
	}
	psr := 1
	if vPosR%2 == 0 {
		psr = 2
	}
	if pRow-vRowL >= vRowR-pRow && vPosR <= sc {
		return vRowR, psr
	} else {
		return vRowL, psl
	}
}
