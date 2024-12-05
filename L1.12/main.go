package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, s := range str {
		set[s] = struct{}{}
	}

	for a := range set {
		fmt.Println(a)
	}
}
