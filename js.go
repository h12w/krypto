// +build js

package main

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

func init() {
	js.Global.Set("JSMain", JSMain)
}

func JSMain(input string) string {
	var ss []string
	kryptoMain(strings.Fields(input), func(answer string) {
		ss = append(ss, answer)
	})
	return strings.Join(ss, "\n")
}
