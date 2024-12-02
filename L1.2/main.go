package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

func main() {
	a := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for num := range a {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}

	wg.Wait()
}
