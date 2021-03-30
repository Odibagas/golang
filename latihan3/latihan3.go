package main

import (
	"fmt"
)

func main() {
	const phi float32 = 3.14
	var r, luas float32
	r = 7
	luas = phi * r * r
	fmt.Println("Luas lingkaran: ", luas)
	var jari, volume float32
	jari = 6
	volume = (4 * phi * jari * jari * jari) / 3
	fmt.Println("Volume bola: ", volume)
}
