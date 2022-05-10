package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang diberikan, buatlah sebuah handler dengan menggunakan HandlerFunc yang menampilkan nama hari, bulan, dan tahun.
// Hint, gunakan time.Weekday, time.Day, time.Month, dan time.Year

func GetHandler() http.HandlerFunc {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		time := time.Now()
		outputTime := fmt.Sprintf("%v, %v %v %v", time.Weekday(), time.Day(), time.Month(), time.Year())
		fmt.Fprint(w, outputTime)

	}
	return handler
}
