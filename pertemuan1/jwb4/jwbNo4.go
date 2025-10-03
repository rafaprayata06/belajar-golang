package main
import "fmt"
func main() {

	var umur = 19
	umur= 21
	fmt.Println(umur)//akan berubah isinya

	const usia= 25
		//klo misal kita ubah isi usia dia akan error
	fmt.Println(usia)
}

