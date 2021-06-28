package main

import (
	"errors"
	"fmt"
)

func sum(scores []int) int {
	res := 0
	for _, score := range scores {
		res += score
	}
	return res
}

func calculate(n1 int, n2 int, opr string) (result float64, err error) {
	switch opr {
	case "+":
		result = float64(n1 + n2)
	case "-":
		result = float64(n1 - n2)
	case "*":
		result = float64(n1 * n2)
	case "/":
		result = float64(n1 / n2)
	default:
		err = errors.New("terjadi error")
	}
	return
}

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)
	result, err := calculate(10, 2, "+")
	fmt.Println(result, err.Error())
	result, err = calculate(10, 2, "-")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "*")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "/")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "=")
	fmt.Println(result, err)
}
