package lab7

import (
	"fmt"
)

type Products interface {
	getName() string
	setPrice(float64)
	getPrice() float64
	changeData(string, float64, string)
	getData()
	applyDiscount(float64)	
}

func calculateDiscount(products []Products) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}

	return sum
}

func RunLab7() {
	product1 := NewBook("Harry Potter", "JK Rowling", 1000.0, "A book on Harry Potter")
	product2 := NewClothes("T-shirt", "Nike", 900.0, "Nice T-shirt")
	product3 := NewElectronics("Lamp", "Diode", 100.0, "Wonderful lamp")

	listOfProduct := []Products{product1, product2, product3}

	fmt.Printf("\nПродукты до изменения\n\n")
	product1.getData()
	product2.getData()
product1.changeData("Improved book Harry Potter", 1500.0, "Interesting book")
product2.changeData("Premium T-shirt", 1000.0, "A high-quality cotton T-shirt.")
	fmt.Printf("\nПродукты после изменения\n\n")
	product1.getData()
	product2.getData()

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt. Printf("%s-------%.2f $\n"， product.getName()，product.getPrice())
	}
	fmt.Printf("Цена корзины до скидки: %.2f $\n\n", calculateDiscount(listOfProduct))

listOfProduct[0].ApplyDiscount(10)
listOfProduct[1].ApplyDiscount(5)

	fmt. Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f $\n"， product.getName(), product.getPrice())
	}
	fmt.Printf("Цена корзины после скидки: %.2f $\n", calculateDiscount(listOfProduct))
}