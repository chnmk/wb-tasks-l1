package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	var str string
	flag.StringVar(&str, "s", "abCdefAaf", "Строка, которую необходимо проверить")
	flag.Parse()

	str = strings.ToLower(str)
	result := true

	// Вариант 1, сложность больше, но без дополнительных структур
	/*
		for i, r1 := range str {
			for j, r2 := range str {
				if i != j && r1 == r2 {
					result = false
				}
			}
		}
	*/

	// Вариант 2, сложность меньше
	strmap := make(map[rune]bool)
	for _, r := range str {
		if strmap[r] {
			result = false
		}
		strmap[r] = true
	}

	fmt.Println(result)
}
