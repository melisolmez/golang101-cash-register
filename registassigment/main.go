package main

import "fmt"

type Item struct {
	name     string
	price    float64
	discount float64
}

func calculatePrice(item Item) float64 {
	if item.discount == 0 {
		return item.price
	}
	var result float64
	result = item.price * item.discount
	item.price -= result
	return item.price
}
func totalPrice(items []Item) float64 {
	var result float64
	for _, item := range items {
		result += item.price
	}
	return result
}

type Describable interface {
	Description() string
}

func (item Item) Description() string {
	if item.discount == 0 {
		return fmt.Sprintf(item.name, item.price)
	}
	var discountedprice float64 = item.price - (item.price * item.discount / 100)
	return fmt.Sprintf(item.name, item.price, discountedprice)

}
func main() {
	var item []Item = []Item{
		{"elma", 10.0, 0},
		{"portakal", 20.0, 0},
	}
	for _, val := range item {
		calculatePrice(val)
	}
}
