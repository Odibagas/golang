package main

import (
	"fmt"
)

func main() {
	var a bool = true
	var b string = "string"
	var c int8 = 1
	var d float64 = 2.5
	var e uint32 = 0
	fmt.Println(a, b, c, d, e)
	f, g := 3, 4
	h := f + g
	fmt.Println(h)
	i := 5
	j := float64(i)
	k := 1.6
	l := j + k
	fmt.Println(l)
}