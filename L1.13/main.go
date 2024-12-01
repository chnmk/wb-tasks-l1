package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	x := 1
	y := 2

	x, y = y, x

	fmt.Printf("x: %d, y: %d", x, y)
}
