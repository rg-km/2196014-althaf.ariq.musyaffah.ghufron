package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Dari contoh sebelumnya tambahkan filter untuk
// menampilkan meja berdasarkan total meja
// key input dari client menggunakan `total`
// contuh URL: http://localhost:8080/table?total=2

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TableHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		total := r.FormValue("total")

		if total != "" {

			result := []Table{}

			for _, table := range data {
				totalStr := strconv.Itoa(table.Total)
				if totalStr == total {
					result = append(result, table)
				}
			}
			resultJSON, err := json.Marshal(result)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// untuk mendaftarkan result sebagai response
			fmt.Println(string(resultJSON))
			w.Write(resultJSON)
			return
		}
		http.Error(w, `{"status":"table not found"}`, http.StatusNotFound)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/table", TableHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
