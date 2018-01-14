package main

import (
	"reflect"
	"testing"
)

func TestPerm(t *testing.T) {
	for _, testcase := range []struct {
		a     []int
		perms [][]int
	}{
		{
			[]int{},
			[][]int{nil},
		},
		{
			[]int{1},
			[][]int{{1}},
		},
		{
			[]int{1, 2},
			[][]int{{1, 2}, {2, 1}},
		},
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	} {
		var perms [][]int
		perm(testcase.a, func(p []int) {
			perms = append(perms, append([]int(nil), p...))
		})
		if !reflect.DeepEqual(perms, testcase.perms) {
			t.Fatalf("expect\n%v\ngot\n%v", testcase.perms, perms)
		}
	}
}
