package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
	Size int
}

func NewStack(size int) Stack {
	return Stack{
		Top:  -1,
		Data: []int{},
		Size: size,
	}
}

func (s *Stack) Push(Elemen int) error {
	//validate if len of data is equal to size means stack is full
	if len(s.Data) == s.Size {
		return errors.New("stack overflow")
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
}

func (s *Stack) Pop() (int, error) {
	//check if stack is empty
	if s.Top == -1 {
		return 0, ErrStackUnderflow
	} else {
		poppedValue := s.Data[s.Top] //get value of top
		s.Top -= 1
		s.Data = s.Data[:len(s.Data)-1] //remove top
		return poppedValue, nil
	}
}
