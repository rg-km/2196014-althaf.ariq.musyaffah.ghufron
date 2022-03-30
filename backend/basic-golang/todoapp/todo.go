package main

import "time"

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	// TODO: answer here
	todos.items = append(todos.items, item)
}

func (todos *Todos) GetAll() []Item {
	return todos.items // TODO: replace this
}

func (todos *Todos) GetUpcoming() []Item {
	//returning items index 1 of 2???
	return todos.items[1:2] // TODO: replace this
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}
