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
		tables := []Table{}
		sumTotal := r.FormValue("total")

		if sumTotal == "" {
			http.Error(w, "invalid total", http.StatusBadRequest)
			return
		}

		total, err := strconv.ParseInt(sumTotal, 16, 32)
		if err != nil {
			http.Error(w, "invalid total", http.StatusBadRequest)
		}

		for _, table := range data {
			if int64(table.Total) == total {
				tables = append(tables, table)
			}
		}
		result, err := json.Marshal(tables)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if len(tables) == 0 {
			http.Error(w, `{"status":"table not found"}`, http.StatusNotFound)
			return
		}

		w.Write(result)
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
