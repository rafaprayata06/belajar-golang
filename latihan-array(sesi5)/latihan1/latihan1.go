package main

import (
	"fmt"
	"strings"
)

type Karyawan struct {
	Nama string
	Usia int
	Gaji float64
}

// Fungsi untuk menambah karyawan
func tambahKaryawan(data *[]Karyawan, nama string, usia int, gaji float64) {
	*data = append(*data, Karyawan{Nama: nama, Usia: usia, Gaji: gaji})
}

// Fungsi untuk menampilkan semua karyawan
func tampilkanKaryawan(data []Karyawan) {
	fmt.Println("\nDaftar Karyawan:")
	for i, k := range data {
		fmt.Printf("%d. Nama: %s | Usia: %d | Gaji: Rp %.2f\n", i+1, k.Nama, k.Usia, k.Gaji)
	}
}

// Fungsi untuk mencari karyawan berdasarkan nama (pakai map)
func cariKaryawanByNama(data []Karyawan, nama string) *Karyawan {
	karyawanMap := make(map[string]Karyawan)
	for _, k := range data {
		karyawanMap[strings.ToLower(k.Nama)] = k
	}

	if k, ada := karyawanMap[strings.ToLower(nama)]; ada {
		return &k
	}
	return nil
}

// Fungsi untuk mengubah data karyawan (pakai pointer)
func ubahGajiKaryawan(k *Karyawan, gajiBaru float64) {
	k.Gaji = gajiBaru
}

// Fungsi untuk menghapus karyawan berdasarkan index
func hapusKaryawan(data *[]Karyawan, index int) {
	if index < 0 || index >= len(*data) {
		fmt.Println("Index tidak valid!")
		return
	}
	*data = append((*data)[:index], (*data)[index+1:]...)
}

func main() {
	var daftarKaryawan []Karyawan

	tambahKaryawan(&daftarKaryawan, "Andi", 25, 5000000)
	tambahKaryawan(&daftarKaryawan, "Budi", 30, 7000000)
	tambahKaryawan(&daftarKaryawan, "Citra", 28, 6500000)

	tampilkanKaryawan(daftarKaryawan)

	fmt.Println("\nCari Karyawan: Budi")
	k := cariKaryawanByNama(daftarKaryawan, "Budi")
	if k != nil {
		fmt.Printf("Ditemukan: %s dengan gaji Rp %.2f\n", k.Nama, k.Gaji)
	}

	fmt.Println("\nUbah gaji karyawan pertama (Andi)...")
	ubahGajiKaryawan(&daftarKaryawan[0], 5500000)
	tampilkanKaryawan(daftarKaryawan)

	fmt.Println("\nHapus karyawan kedua (Budi)...")
	hapusKaryawan(&daftarKaryawan, 1)
	tampilkanKaryawan(daftarKaryawan)
}
