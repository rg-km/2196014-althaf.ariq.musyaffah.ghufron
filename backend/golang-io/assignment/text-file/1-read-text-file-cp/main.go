package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
	// nama text file yang ingin dibaca
	fileName := name

	//membaca text file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return FileData{}, err
	}
	return FileData{
		Name: fileName,
		Size: len(data),
		Data: data,
	}, nil
}
