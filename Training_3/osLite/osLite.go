package oslite

func Oslite(sectCount, partCount int, sects [][2]int) (wo int) {
	m := make(map[int][2]int)
	set := make(map[[2]int]bool)
	for _, s := range sects {
		wo++
		for i := s[0]; i <= s[1]; i++ {
			if v, ok := m[i]; ok {
				if _, ok := set[v]; ok {
					if v == s {
						wo--
					}
					m[i] = s
				} else {
					wo--
					set[v] = true
				}
			} else {
				m[i] = s
			}
		}
	}
	return
}
