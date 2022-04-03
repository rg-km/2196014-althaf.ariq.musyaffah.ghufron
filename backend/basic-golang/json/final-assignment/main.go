package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

// TODO: answer here
type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Item struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	Items []Item `json:"items"`
}
type Ruang struct {
	RuangTamu Items `json:"ruangTamu"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	jsonData, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	return string(jsonData)
}

func NewRuang(r Ruang) Ruang {
	return r
}

func main() {
	itemItem := []Item{
		{
			Nama:   "Meja",
			Jumlah: 20,
			Warna:  "Coklat",
			Ukuran: Ukuran{
				Panjang: "50 cm",
				Tinggi:  "25 cm",
			},
		},
		{
			Nama:   "Kursi",
			Jumlah: 1,
			Warna:  "Hitam",
			Ukuran: Ukuran{
				Panjang: "70 cm",
				Tinggi:  "30 cm",
			},
		},
	}

	items := Items{itemItem}

	ruang := Ruang{
		RuangTamu: items,
	}

	fmt.Println(ruang.EncodeJSON())
}
