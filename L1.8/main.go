package main

import (
	"flag"
	"fmt"
	"strconv"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var i int
	flag.IntVar(&i, "i", 0, "бит, который надо изменить (справа, от наименее значимого, налево)")

	var b uint
	flag.UintVar(&b, "b", 0, "на что изменить бит")
	flag.Parse()

	var myint int64 = 11
	fmt.Printf("Old10: %d\n", myint)
	fmt.Printf("Old2: %s\n", strconv.FormatInt(myint, 2))

	// Используя пакет big
	/*
		I := big.NewInt(myint)
		I.SetBit(I, i, b)

		fmt.Printf("New2: %s\n", strconv.FormatInt(I.Int64(), 2))
		fmt.Printf("New10: %d\n", I.Int64())
	*/

	// Решение, которое используется в пакете big, без импорта пакета
	if b == 1 {
		myint |= (1 << i)
	} else {
		myint &^= (1 << i)
	}

	fmt.Printf("New2: %s\n", strconv.FormatInt(myint, 2))
	fmt.Printf("New10: %d\n", myint)

}
