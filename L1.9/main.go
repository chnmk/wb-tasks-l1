package main

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// TODO: КОНВЕЙЕР

	// Даны два канала.
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	for i := 0; i < 100; i++ {
		// В первый канал пишутся числа из массива.
		ch1 <- i

		// TODO: Последовательно или просто в два канала записывать и с первым ничего не делать?
		number := <-ch1
		ch2 <- number * number

		fmt.Println(<-ch2)
	}
}
