package main

import (
	"math/big"
	"sort"
)

type (
	op int
)

const (
	plus = iota
	minus
	times
	divide
	maxOp
)

var (
	zero big.Rat
)

func krypto(s []int, target int, fn func([]int, []op)) {
	rtarget := big.NewRat(int64(target), 1)
	nums(s, func(a []int) {
		ops(len(s)-1, maxOp, func(op []op) {
			res := calc(a, op)
			if res != nil && res.Cmp(rtarget) == 0 {
				fn(a, op)
			}
		})
	})
}

func calc(a []int, op []op) *big.Rat {
	v := r(a[0])
	for i := 0; i < len(a)-1; i++ {
		if op[i] == divide && a[i+1] == 0 {
			return nil // skip
		}
		v = op[i].calc(v, r(a[i+1]))
	}
	return &v
}

// nums iterates through all the permutations of the numbers
func nums(s []int, fn func([]int)) {
	sort.Ints(s)
	for {
		p := pivot(s)
		if p < 0 {
			break
		}
		np := lastLarger(s, p)
		fn(s)
		s[p], s[np] = s[np], s[p]
		reverse(s[p+1:])
	}
	fn(s)
}
func pivot(s []int) int {
	for i := len(s) - 1; i > 0; i-- {
		if s[i-1] < s[i] {
			return i - 1
		}
	}
	return -1
}
func lastLarger(s []int, p int) int {
	n := s[p]
	for i := len(s) - 1; i > p; i-- {
		if s[i] > n {
			return i
		}
	}
	panic("cannot find last larger")
}
func reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

// ops iterates through all possible operators
func ops(digits int, max op, fn func([]op)) {
	s := make([]op, digits)
	for s[len(s)-1] < max {
		fn(s)
		s[0]++
		for i := 0; i < len(s)-1; i++ {
			if s[i] == max {
				s[i+1]++
				s[i] = 0
			}
		}
	}
}

func (op op) String() string {
	switch op {
	case plus:
		return "+"
	case minus:
		return "-"
	case times:
		return "*"
	case divide:
		return "/"
	}
	return ""
}

func (op op) calc(a, b big.Rat) big.Rat {
	var r big.Rat
	switch op {
	case plus:
		r.Add(&a, &b)
	case minus:
		r.Add(&a, b.Neg(&b))
	case times:
		r.Mul(&a, &b)
	case divide:
		r.Quo(&a, &b)
	}
	return r
}

func r(i int) big.Rat {
	return *big.NewRat(int64(i), 1)
}
