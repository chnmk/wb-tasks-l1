package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).

Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	// Получает количество воркеров через флаг -g
	var numberOfGoroutines int
	flag.IntVar(&numberOfGoroutines, "g", 3, "Количество горутин")
	flag.Parse()

	// Создаёт канал для передачи данных в горутины.
	workCh := make(chan int, numberOfGoroutines)

	// Создаёт контекст для сигнала о завершении работы.
	//
	// Используется завершение работы через контекст, а не через канал,
	// поскольку по умолчанию из канала сигнал о завершении будет прочитан только одной горутиной.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Создаёт горутины для вывода данных из канала workCh.
	var wg sync.WaitGroup
	for i := 0; i < numberOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Завершаю работу горутины")
					return
				case data := <-workCh:
					fmt.Println(data)
				}
			}
		}()
	}

	// Постоянная запись данных до сигнала о завершении.
write:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершаю запись...")
			break write
		default:
			workCh <- rand.IntN(100)
		}
	}

	wg.Wait()
	fmt.Println("Работа завершена")
}
