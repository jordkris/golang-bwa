package main

import "fmt"

func main() {
	// var name string = "Golang"
	var name1 string
	name1 = "Golang 2"
	var flag bool = true
	age := 20

	fmt.Println(name1)
	fmt.Println(flag)
	fmt.Println(age)
	//if else
	if age > 10 {
		fmt.Println("Boleh bermain game")
	} else {
		fmt.Println("Belum boleh bermain game")
	}
	score := 80
	var grade string
	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}
	fmt.Println(grade)
	// switch
	number := 6
	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("DEFAULT")
	}
	// for
	for i := 0; i < 2; i++ {
		fmt.Println("Saya belajar Golang")
	}
	title := "abcde fghij klmno pqrst uvwxyz"
	for index, letter := range title {
		fmt.Println("index: ", index, ", letter: ", string(letter))
	}
	for _, letter := range title {
		fmt.Println("letter: ", string(letter))
	}
}
