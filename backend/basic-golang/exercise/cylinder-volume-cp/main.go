package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var (
		jarijari float32
		tinggi   float32
	)

	fmt.Printf("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&jarijari)
	fmt.Printf("Masukkan tinggi tabung : ")
	fmt.Scan(&tinggi)

	var volume = 3.14 * jarijari * jarijari * tinggi

	fmt.Printf("Volume tabung adalah : %f\n", volume)

}
