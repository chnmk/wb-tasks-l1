package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	// Функция sleep
	time.Sleep(1 * time.Second)

	fmt.Println(1)

	// Вариант 1
	<-time.After(1 * time.Second)

	fmt.Println(2)

	// Вариант 2 - без after
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for range ctx.Done() {
			return
		}

	}()

	wg.Wait()
	fmt.Println(3)
}
