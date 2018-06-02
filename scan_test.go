package main

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	for _, test := range []struct {
		expr   string
		tokens []string
	}{
		{"1", []string{"1"}},
		{"1+2", []string{"1", "+", "2"}},
		{"(1+2)*3", []string{"(", "1", "+", "2", ")", "*", "3"}},
	} {
		tokens := scan(test.expr)
		if !reflect.DeepEqual(tokens, test.tokens) {
			t.Fatalf("expect %v got %v", test.tokens, tokens)
		}
	}
}
