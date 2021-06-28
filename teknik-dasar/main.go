package main

import (
	"fmt"
	"teknik-dasar/calculation"
)

func main() {
	// var name string = "Golang"
	// fmt.Println(name)
	fmt.Println("Halo, belajar golang")
	result1 := calculation.Add(10, 12)
	result2 := calculation.Multiply(10, 12)
	fmt.Println(result1)
	fmt.Println(result2)
}
