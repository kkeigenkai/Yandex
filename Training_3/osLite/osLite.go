package oslite

func Oslite(sectCount, partCount int, sects [][2]int) (wo int) {
	for i := 0; i < len(sects); i++ {
		wo++
		for j := i + 1; j < len(sects); j++ {
			if sects[i][0] <= sects[j][1] && sects[j][0] <= sects[i][1] {
				wo--
				break
			}
		}
	}
	return
}
