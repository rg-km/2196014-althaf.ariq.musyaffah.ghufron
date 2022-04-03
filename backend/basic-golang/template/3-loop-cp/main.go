package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	textTemplate = "{{range . }}" +
		"Peringkat ke-{{ .Rank }}: {{ .Name }}" +
		"{{end}}"

	tmpl, err := template.New("userrank").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, leaderboard.Users)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
