package main

import "fmt"


// Faktor konversi dari kaki ke meter
// 1 kaki = 0.3048 meter
const FAKTOR_KONVERSI = 0.3048

func konversiKakiKeMeter(kaki float64) float64 {
	// Fungsi ini menerima nilai dalam kaki dan mengembalikannya dalam meter.
	return kaki * FAKTOR_KONVERSI
}

func main() {
	fmt.Println("--- Konverter Kaki ke Meter (Go) ---")
	fmt.Print("Masukkan nilai dalam kaki (feet): ")

	var nilaiKaki float64

	// Membaca input angka dari pengguna secara langsung
	_, err := fmt.Scanln(&nilaiKaki)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	// Memanggil fungsi untuk melakukan konversi
	nilaiMeter := konversiKakiKeMeter(nilaiKaki)

	// Menampilkan hasil konversi dengan format dua angka desimal
	fmt.Printf("%.2f kaki sama dengan %.2f meter.\n", nilaiKaki, nilaiMeter)
}