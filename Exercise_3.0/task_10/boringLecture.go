package task10

func BoringLecture(str string) map[rune]int {
	res := make(map[rune]int)
	for i, v := range str {
		pos := i + 1
		res[v] += pos
		res[v] += (len(str) - pos) * pos
	}
	return res
}
