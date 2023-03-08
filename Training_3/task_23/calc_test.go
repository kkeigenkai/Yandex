package task23

import "testing"

type test struct {
	n     int
	expN  int
	expSi []int
}

var tests = []test{
	{1, 0, []int{1}},
	{5, 3, []int{1, 3, 4, 5}},
}

func TestCalc(t *testing.T) {
	for _, test := range tests {
		gotN, gotSi := Calc(test.n)
		if gotN != test.expN || len(gotSi) != len(test.expSi) {
			t.Errorf("( %d ) = %d, %v,\n\twant: %d, %v\n", test.n, gotN, gotSi, test.expN, test.expSi)
			continue
		}
		for i := range gotSi {
			if gotSi[i] != test.expSi[i] {
				t.Errorf("( %d ) = %d, %v,\n\twant: %d, %v\n",
					test.n, gotN, gotSi, test.expN, test.expSi)
				break
			}
		}
	}
}
