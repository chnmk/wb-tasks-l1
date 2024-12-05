package main

import (
	"fmt"

	"math/rand/v2"
)

/*
Вопрос:

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}


Ответ:

Возможны как минимум две проблемы.
	1. Паника если попытаться взять v[:100] от строки меньшего размера
	2. Бессмысленное создание большой переменной v, которая в данном случае не пригодится

Пример улучшенной реализации:
*/

const symbols = "abcdefghijklmnopqrstuvwxyz"

func createHugeString(size int) string {
	// Какая-то логика
	// ...

	b := make([]byte, size)

	for i := range b {
		b[i] = symbols[rand.IntN(len(symbols))]
	}

	return string(b)
}

var justString string

func someFunc() {
	justString = createHugeString(100) // Cразу создаём строку нужного размера, исключая обе проблемы

	// Какая-то логика
	// ...
	fmt.Println(justString)
}

func main() {
	someFunc()
}
