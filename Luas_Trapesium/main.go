package main

import "fmt"

// Luas Trapesium

// Diketahui Trapesium:
// Sisi Atas (a) 	= 5
// Sisi Bawah (b)	= 10
// Tinggi (T)		= 12

func main() {
	var a, b, T uint16 = 5, 10, 12

	Luas := (a + b) * T / 2

	fmt.Print("Jadi luas trapesiumnya adalah: ", Luas)

}
