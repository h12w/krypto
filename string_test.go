package main

import "testing"

func TestString(t *testing.T) {
	for _, test := range []struct {
		nodes nodes
		str   string
	}{
		{},
		{
			[]node{{value: 1}},
			"1",
		},
		{
			[]node{
				{value: plus, children: [2]int8{2, 3}},
				{value: 1},
				{value: 2},
			},
			"1+2",
		},
		{
			[]node{
				{value: plus, children: [2]int8{2, 3}},
				{value: 1},
				{value: plus, children: [2]int8{4, 5}},
				{value: 2},
				{value: 3},
			},
			"1+2+3",
		},
		{
			[]node{
				{value: times, children: [2]int8{2, 3}},
				{value: 1},
				{value: plus, children: [2]int8{4, 5}},
				{value: 2},
				{value: 3},
			},
			"1*(2+3)",
		},
	} {
		str := test.nodes.String()
		if str != test.str {
			t.Fatalf("expect %s got %s", test.str, str)
		}
	}
}
