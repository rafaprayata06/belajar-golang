package main

import "fmt"

func main() {
	// Cetak tabel perkalian 5 dari 1 sampai 10
	for i := 1; i <= 10; i++ {
		hasil := 5 * i
		fmt.Printf("5 x %d = %d\n", i, hasil)
	}
}
