package main

import (
	"fmt"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var num1 float64
	var num2 float64
	var choice = 0
	println("1. rectange area")
	println("2. rectangular area")
	println("3. triangle area")
	println("4. circle area\n")
	fmt.Printf("Pilihan : ")
	fmt.Scan(&choice)

	var hasil float64 = 0

	switch choice {
	case 1:
		fmt.Println("rectang area")
		fmt.Printf("masukkan sisi : ")
		fmt.Scan(&num1)
		hasil = num1 * num1
		fmt.Printf("result = %.0f\n", hasil)
	case 2:
		fmt.Println("rectangular area")
		fmt.Printf("masukkan panjang : ")
		fmt.Scan(&num1)
		fmt.Printf("masukkan lebar : ")
		fmt.Scan(&num2)
		hasil = num1 * num2
		fmt.Printf("result = %.0f\n", hasil)
	case 3:
		fmt.Println("triangle area")
		fmt.Printf("masukkan alas segitiga : ")
		fmt.Scan(&num1)
		fmt.Printf("masukkan tinggi segitiga : ")
		fmt.Scan(&num2)
		hasil = 0.5 * num1 * num2
		fmt.Printf("result = %.2f\n", hasil)
	case 4:
		var pi float64 = 3.14
		println("circle area")
		fmt.Printf("masukkan jari-jari : ")
		fmt.Scan(&num1)
		hasil = pi * num1 * num1
		fmt.Printf("result = %f\n", hasil)

	}
}
