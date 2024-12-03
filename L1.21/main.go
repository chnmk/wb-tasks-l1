package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type resource struct {
	data string
}

func (r resource) get() string {
	return r.data
}

var exampleData = `
<?xml version="1.0" encoding="UTF-8"?>
<example>
	<a>123</a>
	<b>456</b>
</example>
`

var exampleResource = resource{data: exampleData}

func adapter() {

}

func jsonPrinter() {

}

// TODO: доделать, объяснить
func main() {
	d := exampleResource.get()

	jsonPrinter(adapter(d))
}
