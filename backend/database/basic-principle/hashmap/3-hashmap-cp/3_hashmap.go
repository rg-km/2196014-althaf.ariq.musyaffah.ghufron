package main

import "strings"

type PrimaryKey int

type SecondaryKey string

type UserRow struct {
	ID   PrimaryKey   //ID must be unique
	Name SecondaryKey //Name can be non-unique
	Age  int
}

type IndexByID map[PrimaryKey]UserRow

type IndexByName map[SecondaryKey][]PrimaryKey

type UserDB struct {
	ByID   IndexByID
	ByName IndexByName
}

func NewUser() *UserDB {
	return &UserDB{
		ByID:   make(IndexByID),
		ByName: make(IndexByName),
	}
}

func (db *UserDB) Insert(name string, age int) {
	// TODO: answer here
	primaryKey := PrimaryKey(len(db.ByID) + 1)
	secondaryKey := SecondaryKey(name)

	db.ByID[primaryKey] = UserRow{
		ID:   primaryKey,
		Name: secondaryKey,
		Age:  age,
	}

	//save by name
	db.ByName[secondaryKey] = append(db.ByName[secondaryKey], primaryKey)

}

func (db *UserDB) WhereByID(id PrimaryKey) *UserRow {
	m, ok := db.ByID[id]
	if !ok {
		return nil
	}
	return &m
}

func (db *UserDB) WhereByName(name SecondaryKey) []*UserRow {
	ids := db.ByName[name]
	rows := make([]*UserRow, len(ids))
	// TODO: answer here
	for i, id := range ids {
		rows[i] = db.WhereByID(id)
	}

	return rows
}

func (db *UserDB) WhereNameBeginsWith(name string) []*UserRow {
	rows := make([]*UserRow, 0)
	// TODO: answer here
	for row := range db.ByName {
		if strings.HasPrefix(string(row), name) {
			rows = append(rows, db.WhereByName(row)...)
		}
	}

	return rows
}
