package main

import (
	"fmt"
	"strconv"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}
*/

var justString string

func createHugeString(i int) string {
	// Какая-то логика
	// ...

	return strconv.Itoa(i)
}

// TODO
func someFunc() {
	v := createHugeString(1 << 10) // Сдвиг влево на 1 это умножение на 2, сдвиг вправо - деление на 2.
	justString = v[:100]

	fmt.Println(v)
	fmt.Println(justString)
	// runtime error: slice bound out of range [:100] with length 4
	// 2^10 = 1024
}

func main() {
	someFunc()
}
