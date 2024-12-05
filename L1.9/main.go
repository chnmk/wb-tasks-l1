package main

import (
	"fmt"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 151}

	// Даны два канала.
	ch1 := make(chan int, len(arr))
	ch2 := make(chan int, len(arr))

	// В первый канал пишутся числа из массива.
	go func() {
		for _, n := range arr {
			ch1 <- n
		}
		close(ch1)

	}()

	// Во второй - результат операции x*2.
	go func() {
		for i := range ch1 {
			ch2 <- i * i
		}
		close(ch2)
	}()

	// После этого данные из второго канала должны выводиться в stdout.
	for i := range ch2 {
		fmt.Println(i)
	}
}
