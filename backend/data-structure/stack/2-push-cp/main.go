package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
	Size int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack{
		Top:  -1,
		Data: []int{},
		Size: size,
	}
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	//validate if len of data is equal to size means stack is full
	if len(s.Data) == s.Size {
		return ErrStackOverflow
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
}
