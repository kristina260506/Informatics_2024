package main

import (
	"fmt"
	"isuct.ru/informatics2022/lab4"
)
func main() {
	var result float64
	var a float64 = 0.1
	var b float64 = 0.5
	var x [5]float64 = [5]float64{0.2, 0.3, 0.44, 0.6, 0.56}

	fmt.Println("\nзадача A:")
	for i := 0.15; i < 1.37; i += 0.25 {
		result = lab4.Colcul(a, b, i)
		fmt.Println("при x=",i, "y=",result)
	}

	fmt.Println("\nзадача B:")
	for _, i := range x{
		result = lab4.Colcul(a, b, i)
		fmt.Println("при x=",i, "y=",result)
	}
}