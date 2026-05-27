package main

import "fmt"

func main() {
	var nama string
	var umur int
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan umur : ")
	fmt.Scan(&umur)
	fmt.Printf("Halo nama saya %s, saya berumur %d tahun.\n", nama, umur)
}