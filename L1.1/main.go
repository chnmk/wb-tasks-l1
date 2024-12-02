package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).

Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human
}

func (h Human) Study() {
	fmt.Println("Начинаю учиться!")
}

func (h Human) Work() {
	fmt.Println("Начинаю работу!")
}

func (a Action) DoNothing() {
	fmt.Println("Ничего не делаю.")
}

func main() {
	var a Action

	a.Study()
}
