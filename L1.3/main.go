package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.

Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	var result int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	// Создаёт горутины и отправляет в них числа через канал.
	for _, num := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Вычисления производится конкурентно.
			n := <-ch
			square := n * n

			// Потокоезопасная запись в переменную result.
			mu.Lock()
			result += square
			mu.Unlock()
		}()
		// Отправляет число в канал.
		ch <- num
	}

	wg.Wait()
	close(ch)

	fmt.Println(result)
}
