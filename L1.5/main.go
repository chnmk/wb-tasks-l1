package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	// Получает время работы программы в секундах через флаг -s
	var s int
	flag.IntVar(&s, "s", 3, "Количество горутин")
	flag.Parse()

	// Останавливает работу программы через s секунд.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s)*time.Second)
	defer cancel()

	workCh := make(chan int, 1)

write:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершаю работу...")
			break write
		default:
			// Последовательно записывает в канал...
			workCh <- rand.IntN(100)

			// ... и читает из него.
			n := <-workCh
			fmt.Println(n)
		}
	}

	fmt.Println("Работа завершена")
}
