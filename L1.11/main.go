package main

import (
	"fmt"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	// Создание "множества" реализовано в L1.12, здесь оставляю так
	arr1 := []int{10, 25, -1, 8, 0, -35, -45, -10, 2, 21}
	arr2 := []int{11, 26, -1, 8, 0, -34, -45, -15, 1, 22}

	// Есть как минимум два варианта реализации без сортировки: вложенный цикл и мапа

	// Вложенный цикл - самый простой, но менее эффективный вариант
	var result []int
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				result = append(result, i)
			}
		}
	}

	// Вариант через мапу
	var result2 []int
	mymap := make(map[int]int)
	for _, i := range arr1 {
		mymap[i]++
	}
	for _, j := range arr2 {
		mymap[j]++
		if mymap[j] > 1 {
			result2 = append(result2, j)
		}
	}

	fmt.Println(result)
	fmt.Println(result2)
}
