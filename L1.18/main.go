package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type counter struct {
	value int
	mu    sync.Mutex
}

func (c *counter) Add() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	var c counter
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(context context.Context) {
			defer wg.Done()
			for {
				select {
				case <-context.Done():
					return
				default:
					// Какая-то логика
					// ...

					c.Add()
				}
			}
		}(ctx)
	}

	wg.Wait()
	fmt.Println(c.value)
}
