package main

import (
	"fmt"
	"reflect"
)

func main() {

	var productName string = "Laptop"
	var productPrice float64 = 5000.50
	var productQuantity int = 10
	var productAvailable bool = true

	fmt.Println("Product Name: ",
		productName, reflect.TypeOf(productName))

	fmt.Printf("Product Name: %v, Product Price %v ,Product Quantity %v, Product Available %v\n", productName, productPrice, productQuantity, productAvailable)

}
