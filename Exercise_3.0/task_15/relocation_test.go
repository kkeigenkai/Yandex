package task15

import "testing"

type relocationTest struct {
	cnt int
	si  []int
	exp []int
}

var tests = []relocationTest{
	{10, []int{1, 2, 3, 2, 1, 4, 2, 5, 3, 1}, []int{-1, 4, 3, 4, -1, 6, 9, 8, 9, -1}},
}

func TestRelocation(t *testing.T) {
	for _, test := range tests {
		got := Relocation(test.cnt, test.si)
		if len(got) != len(test.exp) {
			t.Errorf("( %d, %v ) = %v,\n\twant: %v\n", test.cnt, test.si, got, test.exp)
		}
		for i := range got {
			if got[i] != test.exp[i] {
				t.Errorf("( %d, %v ) = %v,\n\twant: %v\n", test.cnt, test.si, got, test.exp)
				break
			}
		}
	}
}
