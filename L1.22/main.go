package main

import (
	"flag"
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20.
*/

func main() {
	var a int64
	var b int64
	flag.Int64Var(&a, "a", 4000000000000, "Переменная a")
	flag.Int64Var(&b, "b", -2000000000000, "Переменная b")
	flag.Parse()

	A := big.NewInt(a)
	B := big.NewInt(b)

	var Sum big.Int
	Sum.Add(A, B)

	var Dif big.Int
	var NegB big.Int
	NegB.Neg(B)
	Dif.Add(A, &NegB)

	var Mul big.Int
	Mul.Mul(A, B)

	var Div big.Int
	Div.Div(A, B)

	fmt.Println("Результат сложения: ", Sum.String())
	fmt.Println("Результат вычитания: ", Dif.String())
	fmt.Println("Результат умножения: ", Mul.String())
	fmt.Println("Результат деления: ", Div.String())

}
