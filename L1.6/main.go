package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	var wg sync.WaitGroup

	// Через сигнал отмены в канале
	doneCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-doneCh:
				fmt.Println("Завершаю работу через канал при отмене...")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("жду сигнала")
			}
		}
	}()

	// Через контекст при отмене
	cancelCtx, cancelCancel := context.WithCancel(context.Background())
	defer cancelCancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-cancelCtx.Done():
				fmt.Println("Завершаю работу через контекст при отмене...")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("жду контекст")
			}
		}
	}()

	// Через контекст при истечении времени
	timeoutCtx, cancelTimeout := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelTimeout()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timeoutCtx.Done():
				fmt.Println("Завершаю работу через контекст при истечении времени...")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("жду истечения")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	doneCh <- struct{}{}
	time.Sleep(2 * time.Second)
	cancelCancel()

	wg.Wait()
	fmt.Println("Работа завершена.")
}
