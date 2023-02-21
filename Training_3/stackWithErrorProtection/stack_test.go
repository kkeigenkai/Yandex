package stackwitherrorprotection

import (
	"testing"
)

type stackErrTest struct {
	cmds []string
	exp  []string
}

var tests = []stackErrTest{
	{[]string{
		"push 1",
		"back",
		"exit",
	}, []string{"ok", "1", "bye"}},
}

func TestStackErr(t *testing.T) {
	for _, test := range tests {
		got := StackErr(test.cmds)
		if len(got) != len(test.exp) {
			t.Errorf("( %v ) = %q,\n\twant: %q", test.cmds, got, test.exp)
			continue
		}
		for i, exp := range test.exp {
			if exp != got[i] {
				t.Errorf("( %v ) = %q,\n\twant: %q", test.cmds, got, test.exp)
				break
			}
		}
	}
}
