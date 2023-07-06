package main

import "fmt"

func main(){
	var angka int64

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	if angka > 0 && angka % 7 == 0 {
		fmt.Printf("Angka %d adalah kelipatan 7", angka)
	} else 
	{fmt.Printf("Angka %d bukan merupakan kelipatan 7", angka)}
}