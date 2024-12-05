package main

import (
	"flag"
	"fmt"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func main() {
	var str string
	flag.StringVar(&str, "s", "главрыба", "Строка, которую необходимо перевернуть")
	flag.Parse()

	/*
		// Более короткое, но менее эффективное решение.
		var result1 string
		for _, s := range str {
			result1 = string(s) + result1
		}
	*/

	result := []rune(str)
	max := len(result) - 1

	for i, j := 0, max; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	fmt.Println(string(result))
}
