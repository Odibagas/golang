package main

import (
	"fmt"
)

type rumus struct {
	hasil int
}

func main() {
	var p, l, t int
	fmt.Printf("Masukkan Panjang Balok: ")
	fmt.Scanln(&p)
	fmt.Printf("Masukkan Lebar Balok: ")
	fmt.Scanln(&l)
	fmt.Printf("Masukkan Tinggi Balok: ")
	fmt.Scanln(&t)
	volume := p * l * t
	var a = rumus{volume}
	var b = &a
	fmt.Println(b)
	if p != l {
		fmt.Println("ini pasti bukan kubus")
	} else if p != t {
		fmt.Println("ini pasti bukan kubus")
	} else if l != t {
		fmt.Println("ini pasti bukan kubus")
	} else {
		fmt.Printf("ini bisajadi kubus")
	}
}
