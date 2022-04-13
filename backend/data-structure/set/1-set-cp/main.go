// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

func Intersection(str1, str2 []string) (inter []string) {
	set := make(map[string]bool) //declaring set

	for _, item := range str1 {
		set[item] = true // isi set dengan elemen dari str1 = true
	}

	for _, item := range str2 { //checking if item in str2 is in set
		if _, ok := set[item]; ok {
			inter = append(inter, item)
		}
	}

	inter = RemoveDuplicates(inter) //remove duplicates
	return
}

func RemoveDuplicates(elements []string) (nodups []string) {
	set := map[string]bool{}

	for _, item := range elements {
		if !set[item] {
			// set item true
			set[item] = true
			// Append to nodups slice.
			nodups = append(nodups, item)
		}
	}
	return
}
