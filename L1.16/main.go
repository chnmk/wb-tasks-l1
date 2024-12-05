package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

// Вариант реализации, при котором опорным является последний элемент
func partition(A []int, low, high int) int {
	pivot := A[high]
	i := low

	for j := low; j < high; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}

	A[i], A[high] = A[high], A[i]

	return i
}

func sort(A []int, low, high int) []int {
	if low < high {
		p := partition(A, low, high)
		A = sort(A, low, p-1)
		A = sort(A, p+1, high)
	}

	return A
}

func main() {
	arr := []int{52, 213, 53, 1, 0, -21, 3, 4, 2}

	fmt.Println(sort(arr, 0, len(arr)-1))
}
