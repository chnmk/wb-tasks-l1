package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func checkType(v interface{}) {

	switch v.(type) {

	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Printf("unexpected type")
	}

}

func main() {
	checkType("a")
	checkType(1)
	checkType(true)
	checkType(make(chan int))
}
