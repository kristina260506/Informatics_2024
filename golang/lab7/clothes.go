package lab7

import "fmt"

type Clothes struct {
	name string
	brand string
	price float64
	description string
}

func NewClothes(name string, brand string, price float64, description string) *Clothes {
	clothes := &Clothes{name: name, brand: brand, price: price, description: description}
	return clothes
}

func (c *Clothes) getName() string {
	return c.name
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (c *Clothes) changeData(name string, price float64, description string) {
	c.name = name
	c.price = price
	c.description = description
}

func (c *Clothes) getData() {
	fmt.Printf("name: %s\nbrand: %s\nprice: %.2f\ndescription: %s\n\n", c.name, c.brand, c.price, c.description)
}

func (c *Clothes) applyDiscount(perDiscount float64) {
	c.price = c.price * (1 - (perDiscount / 100))
}