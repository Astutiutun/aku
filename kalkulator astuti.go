package main

import (
	"fmt"
)

func tambah(x, y float64) float64 {
	return x + y
}

func kurang(x, y float64) float64 {
	return x - y
}

func kali(x, y float64) float64 {
	return x * y
}

func bagi(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error: division by zero")
		return 0
	}
	return x / y
}

func main() {
	var pilihan int
	var angka1, angka2 float64
	var lanjutkan bool = true

	fmt.Println("kalkulator sederhana")
	fmt.Println("--------------------")
	for lanjutkan {
		fmt.Println("pilih operasi :")
		fmt.Println("1. penjumlahan")
		fmt.Println("2. pengurangan")
		fmt.Println("3. perkalian")
		fmt.Println("4. pembagian")

		fmt.Print("Masukkan pilihan (1/2/3/4): ")
		fmt.Scanln(&pilihan)

		fmt.Print("Masukkan angka pertama: ")
		fmt.Scanln(&angka1)

		fmt.Print("Masukkan angka kedua: ")
		fmt.Scanln(&angka2)

		switch pilihan {
		case 1:
			fmt.Printf("%.2f + %.2f = %.2f\n", angka1, angka2, tambah(angka1, angka2))
		case 2:
			fmt.Printf("%.2f - %.2f = %.2f\n", angka1, angka2, kurang(angka1, angka2))
		case 3:
			fmt.Printf("%.2f * %.2f = %.2f\n", angka1, angka2, kali(angka1, angka2))
		case 4:
			if angka2 == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Printf("%.2f / %.2f = %.2f\n", angka1, angka2, bagi(angka1, angka2))
			}
		default:
			fmt.Println("pilihan tidak valid")
		}
		var opsi string
		fmt.Print("Lanjutkan operasi? (y/n): ")
		fmt.Scanln(&opsi)

		if opsi != "y" {
			lanjutkan = false
		}
	}
	fmt.Println("Terima kasih telah menggunakan kalkulator ini!")
}
