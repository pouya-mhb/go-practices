package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func createProduct(id string, title string, description string, price float64) *Product {
	productData := Product{
		id:          id,
		title:       title,
		description: description,
		price:       price,
	}
	return &productData
}

func (productData *Product) showPorduct() {
	fmt.Printf("the product info is : %v - %v - %v - %v.", productData.id, productData.title, productData.description, productData.price)
}

func main() {

	product1 := Product{
		id:          "12345",
		title:       "product 1",
		description: "this is a description",
		price:       1.25,
	}

	product2 := createProduct("12345", "product 2", "desc 2", 2.34)

	fmt.Println(product1)
	fmt.Println(product2)

}
