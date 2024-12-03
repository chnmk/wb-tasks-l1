package main

import (
	"flag"
	"fmt"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20.
*/

// 2^20 = 1048576
// TODO: подумать
func main() {
	var a int
	var b int
	flag.IntVar(&a, "a", 4000000000000, "Переменная a")
	flag.IntVar(&b, "b", 2000000000000, "Переменная b")
	flag.Parse()

	fmt.Println(2 << 19)

	a1 := a / (2 << 19)
	b1 := b / (2 << 19)

	fmt.Println(a1)
	fmt.Println(b1)
}
