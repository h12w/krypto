package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("krypto num1 num2 ... numN target")
		return
	}
	kryptoMain(os.Args[1:], func(answer string) {
		fmt.Println(answer)
	})
}

func kryptoMain(args []string, output func(string)) {
	ints, err := parseInts(args)
	if err != nil {
		log.Fatal("wrong arguments", os.Args[1:])
	}
	s, target := ints[:len(ints)-1], ints[len(ints)-1]
	krypto(s, target, func(a []int, op []op) {
		output(format(a, op))
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
