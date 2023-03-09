package task18

import "testing"

type dequeTest struct {
	cmds, exp []string
}

var tests = []dequeTest{
	{
		[]string{
			"push_back 1",
			"back",
			"exit",
		},
		[]string{
			"ok",
			"1",
			"bye",
		},
	},
	{
		[]string{
			"size",
			"push_back 1",
			"size",
			"push_back 2",
			"size",
			"push_front 3",
			"size",
			"exit",
		},
		[]string{
			"0",
			"ok",
			"1",
			"ok",
			"2",
			"ok",
			"3",
			"bye",
		},
	},
	{
		[]string{
			"push_back 1",
			"size",
			"exit",
			"size",
		},
		[]string{
			"ok",
			"1",
			"bye",
		},
	},
}

func TestDuque(t *testing.T) {
	for _, test := range tests {
		got := Deque(test.cmds)
		if len(got) != len(test.exp) {
			t.Errorf("( %v ) = %v,\n\twant: %v\n", test.cmds, got, test.exp)
			continue
		}
		for i := range test.exp {
			if test.exp[i] != got[i] {
				t.Errorf("( %v ) = %v,\n\twant: %v\n", test.cmds, got, test.exp)
				continue
			}
		}
	}
}
