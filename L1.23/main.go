package main

import (
	"flag"
	"fmt"
)

/*
Удалить i-ый элемент из слайса.
*/

func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)

	// Либо можно использовать функцию slices.Delete(), которая работает так же, как и приведённое решение.
}

func main() {
	var a int
	flag.IntVar(&a, "a", 1, "Элемент, который необходимо удалить")
	flag.Parse()

	array := []int{9, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(remove(array, a))
}
