package threeones

var answ = [35]int{}

func ThreeOnes(n int) int {
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 4
	}
	if n == 3 {
		return 7
	}
	if answ[n-1] != 0 {
		return answ[n-1]
	}
	answ[n-1] = ThreeOnes(n-1) + ThreeOnes(n-2) + ThreeOnes(n-3)
	return answ[n-1]
}
