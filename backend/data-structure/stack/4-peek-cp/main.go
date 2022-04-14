package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

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
		return errors.New("stack overflow")
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
}

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
	//check if stack is empty
	if s.Top == -1 {
		return 0, ErrEmptyStack
	} else {
		return s.Data[s.Top], nil
	}
}
