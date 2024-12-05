package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/chnmk/wb-tasks-l1/L1.24/point"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

func main() {
	var ax int
	var ay int
	var bx int
	var by int
	flag.IntVar(&ax, "ax", 3, "Параметр x переменной a")
	flag.IntVar(&ay, "ay", 4, "Параметр y переменной a")
	flag.IntVar(&bx, "bx", 5, "Параметр x переменной b")
	flag.IntVar(&by, "by", 6, "Параметр y переменной b")
	flag.Parse()

	// Можно инкапсулировать параметры структуры в Go, поместив её в отдельный пакет
	var A point.Point
	var B point.Point
	A.SetPoints(ax, ay)
	B.SetPoints(bx, by)

	s1 := B.GetX() - A.GetX()
	s2 := B.GetY() - A.GetY()

	fmt.Println(math.Sqrt(float64(s1*s1 + s2*s2)))
}
