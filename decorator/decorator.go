package main

import "fmt"

type OrderItem interface {
	Price() int64
	Print() string
}

type OrderItemBase struct {
	innerItem OrderItem
	price     int64
	name      string
}

func (i *OrderItemBase) Price() int64 {
	if i.innerItem == nil {
		return i.price
	}
	return i.price + i.innerItem.Price()
}

func (i *OrderItemBase) Print() string {
	if i.innerItem == nil {
		return fmt.Sprintln(i.name)
	}
	return fmt.Sprintln(i.name) + i.innerItem.Print()
}

type Hamburger struct {
	OrderItemBase
}

func NewHamburger(base OrderItem) OrderItem {
	return &Hamburger{
		OrderItemBase{
			innerItem: base,
			price:     16,
			name:      "Hamburger",
		},
	}
}

type ExtraCheese struct {
	OrderItemBase
}

func NewExtraCheese(base OrderItem) OrderItem {
	return &ExtraCheese{
		OrderItemBase{
			innerItem: base,
			price:     3,
			name:      "ExtraCheese",
		},
	}
}

type ExtraLayer struct {
	OrderItemBase
}

func NewExtraLayer(base OrderItem) OrderItem {
	return &ExtraLayer{
		OrderItemBase{
			innerItem: base,
			price:     8,
			name:      "ExtraLayer",
		},
	}
}

type ExtraSauce struct {
	OrderItemBase
}

func NewExtraSauce(base OrderItem) OrderItem {
	return &ExtraSauce{
		OrderItemBase{
			innerItem: base,
			price:     2,
			name:      "ExtraSauce",
		},
	}
}

func main() {
	// base object
	myHamburger := NewHamburger(nil)
	myHamburger = NewExtraCheese(myHamburger) // with extra cheese
	myHamburger = NewExtraLayer(myHamburger)  // with extra layer
	myHamburger = NewExtraSauce(myHamburger)  // with extra sauce
	fmt.Printf("list items : \n %s-----------------\n ",
		myHamburger.Print()) // list of all items combined
	fmt.Printf("Total price: %d", myHamburger.Price()) // aggregation of the items price
}
