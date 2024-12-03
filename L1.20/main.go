package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.

Пример: «snow dog sun — sun dog snow».
*/

func main() {
	var str string
	flag.StringVar(&str, "s", "snow dog sun", "Строка, которую необходимо перевернуть")
	flag.Parse()

	split := strings.Split(str, " ")

	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}

	fmt.Println(split)
}
