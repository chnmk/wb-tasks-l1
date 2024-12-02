package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	arr1 := []int{10, 25, -1, 8, 0, -35, -45, -10, 2, 21}
	arr2 := []int{11, 26, -1, 8, 0, -34, -45, -15, 1, 22}

	// Самый простой вариант
	var result []int
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				result = append(result, i)
			}
		}
	}

	// TODO: ещё можно через мапу и через сортировку. Действительно использовать МНОЖЕСТВО?

	fmt.Println(result)
}
