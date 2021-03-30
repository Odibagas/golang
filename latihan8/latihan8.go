package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 1)
	fmt.Printf("Masukan slice 1: ")
	fmt.Scanln(&slice1[0])
	slice2 := make([]int, 1)
	fmt.Printf("Masukkan slice 2: ")
	fmt.Scanln(&slice2[0])
	var a = slice1[0] + slice2[0]
	slice3 := make([]int, 1)
	slice3 = append(slice3, a)
	fmt.Println("Slice 3:", slice3)
}
