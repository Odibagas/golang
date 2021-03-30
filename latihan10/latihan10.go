package main

import (
	"fmt"
)

type Customer struct {
	Namadep, Namabel, Alamat string
	Umur                     int
}

func (customer Customer) Data() {
	fmt.Println("Nama Depan:", customer.Namadep)
	fmt.Println("Nama Belakang:", customer.Namabel)
	fmt.Println("Umur:", customer.Umur)
	fmt.Println("Alamat:", customer.Alamat)
}

func main() {
	var bagas Customer
	bagas.Namadep = "Bagas"
	bagas.Namabel = "Felix"
	bagas.Umur = 18
	bagas.Alamat = "Bandung"

	bagas.Data()
}
