package main

import "fmt"

func add(number int, numberTwo int) int {
	return number + numberTwo
}

func area_around(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

func area_around2(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

func main() {
	result := add(10, 20)
	fmt.Println(result)
	//multioutput
	luas, keliling := area_around(10, 20)
	fmt.Println(luas, keliling)
	//predefine return value
	luas, keliling = area_around2(30, 40)
	fmt.Println(luas, keliling)
}
