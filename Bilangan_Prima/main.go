package main

import (
	"fmt"
	"math"
)

func main() {

	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	// syarat bilangan prima, bilangan hanya bisa dibagi dengan :
	// 1. Angka 1
	// 2. Bilangan itu sendiri
	if angka <= 1 {
		fmt.Println(angka, "Bukan bilangan prima")
	}

	akar := int(math.Sqrt(float64(angka)))
	fmt.Printf("Akar dari bilangan %d adalah: ", angka)
	fmt.Println(akar)

	for i := 2; i <= akar; i++ {
		if angka%i == 0 {
			fmt.Println(angka, "Bukan merupakan bilangan prima")
			return
		}
	}

	fmt.Println(angka, "Adalah bilangan prima")
	return

}
