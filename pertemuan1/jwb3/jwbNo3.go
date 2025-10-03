package main

import "fmt"

var Global ="saya di pakai dimana saja/global" //bisa di pakai di semua function
	
func main(){
	localVar:="saya lokal"//hanya bisa di akses di funntion klo di akses di luar fungsi dia akan eror yg menyatakan "tidak ada deklarasi dari variable localVar"
	fmt.Println(Global)
	fmt.Println(localVar)
}