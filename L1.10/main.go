package main

import (
	"fmt"
	"strconv"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.

Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	arr := []float64{-25.4, -27.0, -3.1, 3.2, 13.0, 19.0, 15.5, 24.5, 0, -21.0, 32.5}
	result := make(map[string][]float64)

	for _, f := range arr {

		firstDigit := int(f) / 10

		var group string
		if firstDigit == 0 && f < 0 {
			group = "0 to -10"
		} else if firstDigit == 0 && f >= 0 {
			group = "0 to 10"
		} else {
			group = strconv.Itoa(firstDigit * 10)
		}

		result[group] = append(result[group], f)
	}

	fmt.Println(result)
}
