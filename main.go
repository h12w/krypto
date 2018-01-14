package main

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
)

type (
	op int
)

const (
	plus = iota
	minus
	times
	divide
)

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

var (
	zero big.Rat
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("krypto num1 num2 ... numN target")
	}
	ints, err := parseInts(os.Args[1:])
	if err != nil {
		log.Fatal("wrong arguments", os.Args[1:])
	}
	// s := []int{6, 11, 12, 12}
	s, target := ints[:len(ints)-1], ints[len(ints)-1]
	rtarget := big.NewRat(int64(target), 1)
	perm(s, func(a []int) {
		count(len(s)-1, 4, func(op []op) {
			v := r(a[0])
			for i := 0; i < len(a)-1; i++ {
				if op[i] == divide && a[i+1] == 0 {
					return // skip
				}
				v = op[i].calc(v, r(a[i+1]))
			}
			if v.Cmp(rtarget) == 0 {
				fmt.Println(format(a, op))
			}
		})
	})
}

func parseInts(ss []string) ([]int, error) {
	ints := make([]int, len(ss))
	for i := range ints {
		var err error
		ints[i], err = strconv.Atoi(ss[i])
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func format(a []int, op []op) string {
	var buf bytes.Buffer
	v := a[0]
	buf.WriteString(strconv.Itoa(v))
	priority := 1
	for i := 0; i < len(a)-1; i++ {
		switch op[i] {
		case plus, minus:
			priority = 0
		case times, divide:
			if priority == 0 {
				val := "(" + buf.String() + ")"
				buf.Reset()
				buf.WriteString(val)
			}
			priority = 1
		}
		buf.WriteString(op[i].String())
		buf.WriteString(strconv.Itoa(a[i+1]))
	}
	return buf.String()
}

func r(i int) big.Rat {
	return *big.NewRat(int64(i), 1)
}

func count(digits int, max op, fn func([]op)) {
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

func perm(s []int, fn func([]int)) {
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
