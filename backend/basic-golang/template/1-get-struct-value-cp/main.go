package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
	// TODO: answer here
	Name         string
	ScoreAverage float64
}

// main function
func main() {
	buff := new(bytes.Buffer)
	// TODO: answer here
	textTemplate := "Hello {{.Name}}, Nilai rata rata kamu {{.ScoreAverage}}."
	std := Student{"Rogu", 7.8}

	tmp1, err := template.New("tmp1").Parse(textTemplate)
	if err != nil {
		log.Fatal(err)
	}

	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
