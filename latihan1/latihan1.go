package main

import (
	"fmt"
	"time"
)

func swap(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	fmt.Println("Waktu saat ini adalah:", time.Now())
	fmt.Println("Hello", "World!")
	var a, b, c = swap("kata_1", "kata_2", "kata_3")
	fmt.Println(a, b, c)
}
