package main

import (
	"github.com/wangcong15/goassert"
)

type Books struct {
	book_id int
	title   string
}

func main() {
	a := 1
	b := 1 // b := 2
	goassert.AssertEq(a, b)
	_a := 10.2
	_b := 10.2 // _b := 9.2
	goassert.AssertEq(_a, _b)
	c := "hello world"
	d := "hello world" // d := "hello world!"
	goassert.AssertEq(c, d)
	e := true
	f := true // f := false
	goassert.AssertEq(e, f)
	g := [...]int{1, 2, 3}
	h := [...]int{1, 2, 3} // h := [...]int{1, 3, 2}
	goassert.AssertEq(g, h)
	i := &e
	j := &e // j := &g
	goassert.AssertEq(i, j)
	k := Books{1, "Golang"}
	l := Books{1, "Golang"} // l := Books{1, "Python"}
	goassert.AssertEq(k, l)
	m := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	n := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"} // n := map[string]string{"France": "Paris"}
	goassert.AssertEq(m, n)
}
