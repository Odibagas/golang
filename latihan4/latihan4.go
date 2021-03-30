package main

import (
	"fmt"
)

func main() {
	j := 0
	for i := 1; i <= 1000; i++ {
		j += i
	}
	fmt.Println(j)
	k := 0
	l := 0
	for i := 1; i <= 1000000000; i++ {
		k += i
		if k >= 5000000000 {
			l = i
			fmt.Println("loop terakhir:", l)
			break
		}
	}
	fmt.Println("jumlah nilai:", k)
}
