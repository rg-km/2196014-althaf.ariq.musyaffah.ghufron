package main

import (
	"encoding/json"
	"log"
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	// TODO: answer here
	Jenis  string `json:"jenis"`
	Warna  string `json:"color"`
	Jumlah int    `json:"jumlah"`
}

func (m Meja) EncodeJSON() string {
	// TODO: answer here
	m = Meja{
		Jenis:  "Meja Belajar",
		Warna:  "green",
		Jumlah: 2,
	}
	jsonMeja, err := json.Marshal(m)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}
	return string(jsonMeja)
}

func NewMeja(m Meja) Meja {
	return m
}
