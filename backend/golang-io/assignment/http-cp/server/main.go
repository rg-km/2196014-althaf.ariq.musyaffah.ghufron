package main

import (
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, world!"))
			return
		}

	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Hello, world!"}`))
			return
		}

	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
