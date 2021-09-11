package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
)

type CoffeeType int

const (
	Latte CoffeeType = iota
	Espresso
	Americano
	Mocha
)

type Coffee struct {
	Type            CoffeeType
	BasePrice       int
	DiscountPercent int
	Tax             int
}

func (c *Coffee) Print() {
	fmt.Println("------Coffee------")
	fmt.Println("type :" + strconv.Itoa(int(c.Type)))
	fmt.Println("price :" + strconv.Itoa(c.BasePrice))
	fmt.Println("discount percent :" + strconv.Itoa(c.DiscountPercent))
	fmt.Println("tax :" + strconv.Itoa(c.Tax))
}

func (c *Coffee) deepCopy() *Coffee {
	b := bytes.Buffer{}
	_ = gob.NewEncoder(&b).Encode(c)

	result := Coffee{}
	_ = gob.NewDecoder(&b).Decode(&result)
	return &result
}

var coffeeContainer = map[CoffeeType]*Coffee{
	Americano: {Americano, 15, 3, 15},
	Latte:     {Latte, 20, 4, 15},
	Espresso:  {Espresso, 12, 0, 15},
	Mocha:     {Mocha, 22, 4, 15},
}

func NewCoffee(coffeeType CoffeeType) *Coffee {
	coffee := coffeeContainer[coffeeType]
	if coffee == nil {
		panic("cannot find the coffee type in the container")
	}
	return coffee.deepCopy()
}

func main() {
	latte := NewCoffee(Latte)
	latte.Print()

	espresso := NewCoffee(Espresso)
	espresso.Print()
}
