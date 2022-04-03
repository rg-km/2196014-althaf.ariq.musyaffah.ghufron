package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan condition pada template.
// Lengkapi function ExecuteToByteBuffer dan variabel textTemplate sehingga memiliki keluaran:
// 1. Jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
// 2. Jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	textTemplate = "{{if (eq .Balance 0)}}" +
		"Akun {{.Name}} dengan Nomor {{.Number}} tidak memiliki saldo." +
		"{{ else }}Akun {{.Name}} dengan Nomor {{.Number}} memiliki saldo sebesar ${{.Balance}}." +
		"{{ end }}"

	tmpl, err := template.New("account").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, account)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

// func main() {
// 	account := Account{
// 		Name:    "Marwan",
// 		Number:  "12345",
// 		Balance: 10000,
// 	}

// 	res, _ := ExecuteToByteBuffer(account)
// 	fmt.Println(string(res))
// }
