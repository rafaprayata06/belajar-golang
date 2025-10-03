package main

import "fmt"

func main() {
    var fahrenheit float64
    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&fahrenheit)

    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Printf("%.2f °F = %.2f °C\n", fahrenheit, celsius)
}
