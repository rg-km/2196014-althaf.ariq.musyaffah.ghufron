package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan http.HandleFunc

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	time := time.Now()
	fmt.Fprintf(writer, "%v, %v %v %v", time.Weekday(), time.Day(), time.Month(), time.Year())
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello there")
	} else {
		fmt.Fprintf(w, "Hello, %v!", name)
	}

}

func main() {
	http.HandleFunc("/time", TimeHandler)
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
