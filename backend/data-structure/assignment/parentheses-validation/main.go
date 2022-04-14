package main

import "github.com/ruang-guru/playground/backend/data-structure/assignment/parentheses-validation/stack"

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// validasi panjang stringnya tdk genap (%) atau string kosong -> return false
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	//import struct stack and initialize stack and fill attributes
	stack := new(stack.Stack)
	stack.Top = -1
	stack.Data = []rune{}

	//make map contains all parentheses
	parentheses := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	//iterate string
	for _, char := range s {
		//if char is exists, push to stack
		if _, ok := parentheses[char]; ok {
			stack.Push(char)
		} else {
			//if char is not exists, check if stack is empty
			if stack.Top == -1 {
				return false
			}

			//if stack is not empty, check if char is equal to top of stack
			if char != parentheses[stack.Data[stack.Top]] {
				return false
			}

			//if char is equal to top of stack, pop the top of stack
			stack.Pop()
		}

	}

	return stack.IsEmpty()

}
