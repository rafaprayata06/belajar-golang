package main
import (
"fmt"
)
type Mahasiswa struct {

NIM string
Nama string
Usia int
Jurusan string
}
func tambahMahasiswa(data *[]Mahasiswa, mhs Mahasiswa) {
*data = append(*data, mhs)
}
func tampilkanMahasiswa(data []Mahasiswa) {
for _, mhs := range data {
fmt.Printf("NIM: %s, Nama: %s, Usia: %d, Jurusan: %s\n", mhs.NIM, mhs.Nama, mhs.Usia,
mhs.Jurusan)
}
}
func cariMahasiswa(data []Mahasiswa, nim string) *Mahasiswa {
for i, mhs := range data {
if mhs.NIM == nim {
return &data[i]
}
}
return nil
}
func hapusMahasiswa(data *[]Mahasiswa, nim string) bool {
for i, mhs := range *data {
if mhs.NIM == nim {
*data = append((*data)[:i], (*data)[i+1:]...)
return true
}
}
return false
}
func filterMahasiswa(data []Mahasiswa, filterFunc func(Mahasiswa) bool) []Mahasiswa {
var filtered []Mahasiswa
for _, mhs := range data {
if filterFunc(mhs) {
filtered = append(filtered, mhs)
}
}
return filtered
}
func main() {

var dataMahasiswa []Mahasiswa
// Menambahkan data mahasiswa
tambahMahasiswa(&dataMahasiswa, Mahasiswa{
NIM: "001",
Nama: "Andi",
Usia: 20,
Jurusan: "Informatika",
})
tambahMahasiswa(&dataMahasiswa, Mahasiswa{
NIM: "002",
Nama: "Budi",
Usia: 21,
Jurusan: "Sistem Informasi",
})
tambahMahasiswa(&dataMahasiswa, Mahasiswa{
NIM: "003",
Nama: "Cici",
Usia: 19,
Jurusan: "Informatika",
})
// Menampilkan semua data mahasiswa
fmt.Println("=== Data Mahasiswa ===")
tampilkanMahasiswa(dataMahasiswa)
// Mencari mahasiswa berdasarkan NIM
nimDicari := "002"
mhs := cariMahasiswa(dataMahasiswa, nimDicari)
if mhs != nil {
fmt.Printf("\nMahasiswa dengan NIM %s ditemukan:\n", nimDicari)
fmt.Printf("Nama: %s, Usia: %d, Jurusan: %s\n", mhs.Nama, mhs.Usia, mhs.Jurusan)
} else {
fmt.Printf("\nMahasiswa dengan NIM %s tidak ditemukan.\n", nimDicari)
}
// Menghapus mahasiswa berdasarkan NIM
nimDihapus := "001"
if hapusMahasiswa(&dataMahasiswa, nimDihapus) {
fmt.Printf("\nMahasiswa dengan NIM %s berhasil dihapus.\n", nimDihapus)
} else {
fmt.Printf("\nMahasiswa dengan NIM %s tidak ditemukan.\n", nimDihapus)
}
// Menampilkan data mahasiswa setelah penghapusan
fmt.Println("\n=== Data Mahasiswa Setelah Penghapusan ===")
tampilkanMahasiswa(dataMahasiswa)

// Memfilter mahasiswa jurusan Informatika menggunakan fungsi anonim
fmt.Println("\n=== Mahasiswa Jurusan Informatika ===")
mhsInformatika := filterMahasiswa(dataMahasiswa, func(mhs Mahasiswa) bool {
return mhs.Jurusan == "Informatika"
})
tampilkanMahasiswa(mhsInformatika)
}