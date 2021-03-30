package main

import (
	"fmt"
)

func main() {
	var alas, tinggi, luas float32
	fmt.Printf("Masukkan Alas: ")
	fmt.Scan(&alas)
	fmt.Printf("Masukkan Tinggi: ")
	fmt.Scan(&tinggi)
	luas = 0.5 * alas * tinggi
	if alas > tinggi {
		fmt.Println("luas segitiganya adalah :", luas, ";Segitiga-nya Lebar")
	} else if alas < tinggi {
		fmt.Println("luas segitiganya adalah :", luas, ";Segitiga-nya Tinggi")
	} else {
		fmt.Println("luas segitiganya adalah :", luas, ";Segitiga-nya sama kaki")
	}

	j := 1
	k := 0
	for i := 1; i <= 10000000000; i++ {
		if i%2 == 0 {
			j += 1
		} else {
			k += 1
		}
	}
	fmt.Println("Jumlah bilangan genap:", j)
	fmt.Println("Jumlah bilangan ganjil:", k)
}
