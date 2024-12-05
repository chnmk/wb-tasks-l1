package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Данные в формате XML
var exampleData = `
<?xml version="1.0" encoding="UTF-8"?>
<example>
	<a>123</a>
	<b>456</b>
</example>
`

// Абстрактный источник данных
type resource struct {
	data string
}

// Конкретный источник данных
var exampleResource = resource{data: exampleData}

// Функция для получения данных из источника
func (r resource) get() string {
	return r.data
}

// ===================================

// Структура для преобразования данных
type exampleStruct struct {
	XMLName xml.Name `xml:"example"`
	A       int      `xml:"a"`
	B       int      `xml:"b"`
}

// Адаптер, преобразовывающий данные в нужный формат
func adapter(d string) string {
	var example exampleStruct
	err := xml.Unmarshal([]byte(d), &example)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	json, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	return string(json)
}

func main() {
	// Получает данные
	d := exampleResource.get()

	// Печатает преобразованные через адаптер данные
	fmt.Println(adapter(d))
}
