// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "keen"
	var str2 = "knee"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {

	if len(str1) != len(str2) {
		return "Bukan Anagram"
	}

	strMap := make(map[string]int)

	for _, val := range str1 {
		x := strMap[string(val)]

		if x == 0 {
			strMap[string(val)] = 1
		} else {
			strMap[string(val)] = x + 1
		}
	}

	for _, val := range str2 {
		x := strMap[string(val)]

		if x == 0 {
			strMap[string(val)] = 1
		} else {
			strMap[string(val)] = x + 1
		}
	}

	isAnagram := true

	fmt.Println(strMap, isAnagram)

	for _, val := range strMap {
		if val%2 != 0 {
			isAnagram = false
			break
		}
	}

	if isAnagram {
		return "Anagram"
	}
	return "Bukan Anagram"

	// str1Map := map[rune]int{}
	// str2Map := map[rune]int{}

	// for _, val := range str1 {
	// 	str1Map[val]++
	// }

	// for _, val := range str2 {
	// 	str2Map[val]++
	// }

	// for key, val := range str1Map {
	// 	if str2Map[key] != val {
	// 		return "Bukan Anagram"
	// 	}
	// }

	// return "Anagram"
}
