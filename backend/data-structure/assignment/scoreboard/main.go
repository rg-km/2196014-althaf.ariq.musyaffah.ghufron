package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	score1 := s[i].Correct*4 - s[i].Wrong
	score2 := s[j].Correct*4 - s[j].Wrong

	if score1 > score2 {
		return true
	} else if score1 < score2 {
		return false
	}

	if s[i].Correct > s[j].Correct {
		return true
	} else if s[i].Correct < s[j].Correct {
		return false
	}

	return s[i].Name < s[j].Name
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] //swap the element
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},  //score: 3*4 - 2 = 10
		{"Agus", 3, 4, 0},  //score: 3*4 - 4 = 8
		{"Doni", 3, 0, 7},  //score: 3*4 - 0 = 12
		{"Ega", 3, 0, 7},   //score: 3*4 - 0 = 12
		{"Anton", 2, 0, 5}, //score: 2*4 - 0 = 8
	})
	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
	//expected output: ["Doni", "Ega", "Levi", "Agus", "Anton"]
}
