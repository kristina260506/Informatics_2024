package lab7

import "fmt"

type Electronics struct {
	name string
	model string
	price float64
	description string
}

func NewElectronics(name string, model string, price float64, description string) *Electronics {
	electronics := &Electronics{name: name, model: model, price: price, description: description}
	return electronics
}

func (e *Electronics) getName() string {
	return e.name
}

func (e *Electronics) setPrice(newPrice float64) {
	e.price = newPrice
}

func (e *Electronics) getPrice() float64 {
	return e.price
}

func (e *Electronics) changeData(name string, price float64, description string) {
	e.name = name
	e.price = price
	e.description = description
}

func (e *Electronics) getData() {
	fmt.Printf("name: %s\nmodel: %s\nprice: %.2f\ndescription: %s\n\n", e.name, e.electronics, e.price, e.description)
}

func (e *Electronics) applyDiscount(perDiscount float64) {
	e.price = e.price * (1 - (perDiscount / 100))
}