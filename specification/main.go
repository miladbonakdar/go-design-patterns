package main

import "fmt"

type Category int

const (
	Cat1 Category = iota
	Cat2
	Cat3
)

type Product struct {
	name     string
	price    uint
	category Category
}

type Specification interface {
	Satisfied(order Product) bool
}

type PriceSuggestion struct {
	fromPrice uint
	toPrice   uint
}

func (p PriceSuggestion) Satisfied(order Product) bool {
	return order.price > p.fromPrice && order.price <= p.toPrice
}

type CategorySuggestion struct {
	category Category
}

func (p CategorySuggestion) Satisfied(order Product) bool {
	return order.category == p.category
}

type AndSuggestions struct {
	first  Specification
	second Specification
}

func (p AndSuggestions) Satisfied(order Product) bool {
	return p.first.Satisfied(order) && p.second.Satisfied(order)
}

type OrSuggestions struct {
	first  Specification
	second Specification
}

func (p OrSuggestions) Satisfied(order Product) bool {
	return p.first.Satisfied(order) || p.second.Satisfied(order)
}

func GetSuggestions(input []Product, spec Specification) []Product {
	items := make([]Product, 0)
	for _, p := range input {
		if spec.Satisfied(p) {
			items = append(items, p)
		}
	}
	return items
}

func main() {
	products := []Product{
		{name: "one", price: 10, category: Cat1},
		{name: "two", price: 20, category: Cat2},
		{name: "three", price: 30, category: Cat3},
		{name: "four", price: 40, category: Cat1},
		{name: "five", price: 50, category: Cat2},
		{name: "six", price: 60, category: Cat3},
		{name: "seven", price: 70, category: Cat1},
		{name: "eight", price: 80, category: Cat2},
		{name: "nine", price: 90, category: Cat3},
	}
	priceSuggestion := &PriceSuggestion{fromPrice: 20, toPrice: 60}
	suggestions := GetSuggestions(products, priceSuggestion)
	fmt.Println("------ suggestions by price ------")
	PrintProducts(suggestions)

	cat2Suggestion := &CategorySuggestion{category: Cat2}
	suggestions = GetSuggestions(products, cat2Suggestion)
	fmt.Println("------ suggestions by category ------")
	PrintProducts(suggestions)

	priceAndCategorySuggestion := &AndSuggestions{
		first:  priceSuggestion,
		second: cat2Suggestion,
	}
	suggestions = GetSuggestions(products, priceAndCategorySuggestion)
	fmt.Println("------ suggestions by category and price ------")
	PrintProducts(suggestions)

	priceOrCategorySuggestion := &OrSuggestions{
		first:  priceSuggestion,
		second: cat2Suggestion,
	}
	suggestions = GetSuggestions(products, priceOrCategorySuggestion)
	fmt.Println("------ suggestions by category or price ------")
	PrintProducts(suggestions)
}

func PrintProducts(items []Product) {
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		fmt.Print(item.name + ",")
	}
	fmt.Println("\n----------------------------")
}
