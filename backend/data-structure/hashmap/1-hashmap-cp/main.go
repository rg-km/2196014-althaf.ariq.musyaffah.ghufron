package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	AgeDist := make(map[int]int)
	for _, person := range people {
		AgeDist[person.age]++
	}
	return AgeDist
}

func FilterByAge(people []Person, age int) []Person {
	result := make(map[int][]Person)
	for _, person := range people {
		result[person.age] = append(result[person.age], person)
	}
	return result[age]
}
