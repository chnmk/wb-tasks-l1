package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	time.Sleep(1 * time.Second)

	fmt.Println(1)

	<-time.After(1 * time.Second)

	fmt.Println(2)

	// TODO
	c := make(chan time.Time, 1)
	_ = c

}
