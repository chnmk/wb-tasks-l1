package main

import "flag"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func main() {
	var str string
	flag.StringVar(&str, "s", "главрыба", "Строка, которую необходимо перевернуть")
	flag.Parse()
}
