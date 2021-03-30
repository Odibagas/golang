package main

import "fmt"

func main() {
	var array1 [1]int
	fmt.Printf("Masukan array pertama: ")
	fmt.Scanln(&array1[0])
	var array2 [1]int
	fmt.Printf("Masukkan array kedua: ")
	fmt.Scanln(&array2[0])
	a := array1[0] + array2[0]
	fmt.Println(array1[0], "+", array2[0], "adalah", a)
	var array3 [1]int
	array3[0] = a
	fmt.Println("Nilai Array setelah data penjumlahan tersimpan: ", array3)
}
