package main

import "fmt"

func main() {
	title := "Golang the best language"
	for index, letter := range title {
		if index%2 == 1 {
			fmt.Println("letter :", string(letter))
		}
	}
	for _, letter := range title {
		switch string(letter) {
		case "a", "i", "u", "e", "o":
			fmt.Println("letter :", string(letter))
		}
	}
}
