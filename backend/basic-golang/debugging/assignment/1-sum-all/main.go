package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{5, 0, 1, 2, 3}

	res := SumAll(arr)
	fmt.Println(res)

	// Try correct answer:
	resCorrect := SumAllCorrect(arr)
	fmt.Println(resCorrect)
}

func SumAll(arr []int) int {
	res := 0
	for val := range arr {
		log.Println(val)
		res += val
	}
	return res
}

func SumAllCorrect(arr []int) int {
	res := 0
	for _, val := range arr {
		log.Println(val)
		res += val
	}
	return res
}
