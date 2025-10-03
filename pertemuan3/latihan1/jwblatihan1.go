package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {

		if i == 5 {
			continue // kalau i == 5, lewati iterasi ini
		}

		fmt.Println(i) // baris ini tidak akan jalan saat i == 5
	}
}
