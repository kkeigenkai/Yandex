package task29

import (
	"reflect"
	"testing"
)

type args struct {
	cnt int
	si  []int
}

type test struct {
	name       string
	args       args
	wantAmount int
	wantCoup   [2]int
	wantDay    int
}

var tests = []test{
	{"base test", args{5, []int{35, 40, 101, 59, 63}}, 235, [2]int{0, 1}, 5},
}

func TestCafe(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAmount, gotCoup, gotDay := Cafe(tt.args.cnt, tt.args.si)
			if gotAmount != tt.wantAmount {
				t.Errorf("Cafe() gotAmount = %v, want %v", gotAmount, tt.wantAmount)
			}
			if !reflect.DeepEqual(gotCoup, tt.wantCoup) {
				t.Errorf("Cafe() gotCoup = %v, want %v", gotCoup, tt.wantCoup)
			}
			if gotDay != tt.wantDay {
				t.Errorf("Cafe() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}
