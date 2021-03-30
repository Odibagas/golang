package main

import (
	"fmt"
)

func main() {

	angka := map[int]int{
		1: 2,
		2: 3,
		3: 4,
		4: 5,
		5: 6,
	}
	fmt.Println("before deletion:", angka)
	delete(angka, 3)
	fmt.Println("after deletion:", angka)
}
