package main

import (
	"github.com/wangcong15/goassert"
)

func main() {
	a := 2
	b := 1 // b := 2
	goassert.AssertGt(a, b)
	_a := 10.2
	_b := 10.1 // _b := 10.2
	goassert.AssertGt(_a, _b)
}
