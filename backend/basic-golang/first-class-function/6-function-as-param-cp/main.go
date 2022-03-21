package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
	// TODO: answer here
	printString := func(fungsi func() string) {
		result := fungsi()
		fmt.Println(result)
	}

	goodNight := func() string {
		return "selamat malam"
	}

	printString(goodNight)

}
