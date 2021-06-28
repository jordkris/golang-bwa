package main

import (
	"fmt"
	"strconv"
)

func main() {
	// array
	var languages [5]string
	for i := 0; i < len(languages); i++ {
		languages[i] = "golang " + strconv.Itoa(i)
	}
	fmt.Println(languages)
	languages2 := [...]string{
		"Go",
		"Java",
		"Python",
		"HTML",
	}
	for index, letter := range languages2 {
		fmt.Println("index: ", index, ", letter: ", letter)
	}
	var numbers [5]int
	numbers2 := [...]int{5, 4, 3, 2, 1}
	n := 0
	for n < 5 {
		numbers2[n] = n
		numbers[n] = numbers2[n] + 1
		n++
	}
	fmt.Println(numbers, numbers2)
	var flags [5]bool
	fmt.Println(flags)

	//slice
	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "Playstation 4")
	gamingConsoles = append(gamingConsoles, "Playstation 3")
	odd := []int{1, 3, 5, 7, 9}
	fmt.Println(gamingConsoles, odd)
	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

	//map        [  key ]value
	myMap := map[string]int{}
	myMap["ruby"] = 9
	myMap["javascript"] = 8
	myMap2 := map[int]string{0: "kucing", 2: "anjing", 4: "ayam"}
	fmt.Println(myMap, myMap["ruby"], myMap2, myMap2[1], myMap2[2])
	fmt.Println(" ")
	//map (advanced)
	myMap3 := map[string]string{
		"ruby": "is beautiful",
		"go":   "is super fast",
	}
	for key, value := range myMap3 {
		fmt.Println("key:", key, "value:", value)
	}
	fmt.Println("=====")
	delete(myMap3, "ruby")
	for key, value := range myMap3 {
		fmt.Println("key:", key, "value:", value)
	}
	fmt.Println(" ")
	val := myMap3["go"]
	value, isAvailable := myMap3["python"]
	fmt.Println(val)
	fmt.Println(value, isAvailable)

	// slice x map
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Mario", "score": "E"},
	}
	fmt.Println(students)
	for _, student := range students {
		fmt.Println(student)
		fmt.Println(student["name"], "-", student["score"])
	}
}
