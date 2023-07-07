package main

import "fmt"

// Luas Trapesium

// Diketahui Trapesium:
// Sisi Atas (a) 	= 5
// Sisi Bawah (b)	= 10
// Tinggi (T)		= 12

func main() {
	var a, b, T int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai T: ")
	fmt.Scanln(&T)

	Luas := float32(a + b) * float32(T) / 2

	fmt.Print("Jadi luas trapesiumnya adalah: ", Luas)

}
