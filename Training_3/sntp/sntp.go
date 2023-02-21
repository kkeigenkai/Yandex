package sntp

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SNTP(ta, tb, tc string) string {
	timeA := getSec(strToInt(strings.Split(ta, ":")))
	timeB := getSec(strToInt(strings.Split(tb, ":")))
	timeC := getSec(strToInt(strings.Split(tc, ":")))

	delay := int(math.Round(float64(timeC-timeA) / 2))
	if timeA > timeC {
		delay = int(math.Round(float64(timeC+24*3600-timeA) / 2))
	}

	answ := timeB + delay
	h := answ / 3600
	if h >= 24 {
		h -= 24
	}
	m := answ / 60 % 60
	s := answ % 60

	return fmt.Sprintf("%s:%s:%s", getStr(h), getStr(m), getStr(s))
}

func strToInt(ss []string) []int {
	si := make([]int, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		si = append(si, i)
	}
	return si
}

func getSec(si []int) int {
	return si[0]*3600 + si[1]*60 + si[2]
}

func getStr(t int) string {
	ts := fmt.Sprintf("%d", t)
	if t < 10 {
		ts = "0" + ts
	}
	return ts
}
