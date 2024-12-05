package main

import (
	"errors"
	"fmt"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func search(A []int, target int) (int, error) {
	r := len(A) - 1

	for l := 0; l < r; l++ {
		m := int(l + r/2)

		if A[m] < target {
			l = m + 1
		} else if A[m] > target {
			r = m - 1
		} else {
			return m, nil
		}
	}

	return 0, errors.New("value not found")
}

func main() {
	arr := []int{1, 2, 4, -1, 20, 5, -333, 0}

	result, err := search(arr, -1)
	if err != nil {
		fmt.Println("Число не найдено:", err)
	} else {
		fmt.Println("Положение числа -1 в массиве:", result)
	}
}
