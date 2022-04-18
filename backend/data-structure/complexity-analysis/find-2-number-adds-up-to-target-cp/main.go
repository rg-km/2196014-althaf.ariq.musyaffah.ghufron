// Temukan dua angka di array yang menghasilkan target dalam Go (Golang)
// Sebagai contoh, kita memiliki sebuah array: [2, 5, 1, 3]
// Targetnya adalah 4
// Lalu outputnya adalah index: [2, 3]

// Penjelasan
// Nomor 1 di index 2
// Nomor 3 di index 3
// dan 1+3 = 4

// Perhatikan bahwa array tidak disortir

// Expektasi Time Complexity adalah -> O(n)

// Kita bisa menggunakan hash map untuk solusinya. Hal ini didasarkan pada gagasan bahwa

// Jika salah satu bilangan tersebut adalah x
// Kemudian nomor lainnya akan menjadi target-x
// Jadi jika untuk angka x kita periksa bahwa target-x ada di hash. Jika ya maka kita tahu kita punya solusinya

package main

import "fmt"

func main() {
	output := TwoTargetSums([]int{2, 5, 1, 3}, 6)
	fmt.Println(output)
}

func TwoTargetSums(nums []int, target int) []int {
	numberMap := make(map[int]int) // key: number, value: index
	output := make([]int, 2)       // index of the two numbers
	for i := 0; i < len(nums); i++ {
		val, ok := numberMap[target-nums[i]] // check if target-nums[i] is in the map
		// TODO: answer here
		if ok { // if target - nums[i] in map
			output[0] = val
			output[1] = i
			return output
		}
		numberMap[nums[i]] = i // else, add nums[i] to map
	}
	return output
}
