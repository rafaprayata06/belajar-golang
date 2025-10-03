package main

import "fmt"

// Fungsi untuk cek apakah bilangan prima
func isPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ { // cek pembagi dari 2 sampai akar n
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Bilangan prima dari 1 sampai 50:")

	for i := 1; i <= 50; i++ {
		if isPrima(i) {
			fmt.Print(i, " ")
		}
	}
}
