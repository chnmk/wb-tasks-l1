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
	// TODO

	var b int64
	flag.Int64Var(&b, "b", 2, "бит, который надо изменить")

	var v int64
	flag.Int64Var(&v, "v", 1, "на что изменить бит")
	flag.Parse()

	var i int64 = 11
	fmt.Printf("Old10: %d\n", i)
	fmt.Printf("Old2: %s\n", strconv.FormatInt(i, 2))

	i |= (v << b)
	fmt.Printf("New2: %s\n", strconv.FormatInt(i, 2))
	fmt.Printf("New10: %d\n", i)
}
