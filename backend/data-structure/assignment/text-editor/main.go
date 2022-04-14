package main

import (
	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	te.RedoStack.SetToEmpty()
	te.UndoStack.Push(ch)
}

func (te *TextEditor) Read() []rune {
	//initiate a new stack and call newStack function
	tmp := stack.NewStack()
	result := []rune{}

	//when text editor has no inserted character
	if te.UndoStack.Top == -1 {
		return nil
	}

	for te.UndoStack.Top > -1 {
		//call pop function from stack
		pop, _ := te.UndoStack.Pop()
		//push the popped element to the new stack
		tmp.Push(pop)
	}

	for tmp.Top > -1 {
		//call peek function from stack
		peek, _ := tmp.Peek()

		//push the popped element to the new stack
		result = append(result, peek)
		tmp.Pop()

		//push the popped element to the new stack
		te.UndoStack.Push(peek)
	}
	return result
}

func (te *TextEditor) Undo() {
	//call peek function to get the top element
	peek, err := te.UndoStack.Peek()
	if err != nil {
		return
	}

	//call pop function to remove the top element
	te.UndoStack.Pop()

	te.RedoStack.Push(peek)
}

func (te *TextEditor) Redo() {
	peek, err := te.RedoStack.Peek()
	if err != nil {
		return
	}

	te.RedoStack.Pop()

	te.UndoStack.Push(peek)
}
