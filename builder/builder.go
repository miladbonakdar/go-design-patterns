package main

import (
	"fmt"
)

//this is our complex object
type Sword struct {
	damage int
	length int
	weight int
	price  int
}

type SwordBuilder struct {
	sword *Sword
}

func NewSwordBuilder() *SwordBuilder {
	return &SwordBuilder{
		// you can set sward default values here
		sword: &Sword{
			damage: 0,
			length: 0,
			weight: 0,
			price:  0,
		},
	}
}

func (b *SwordBuilder) SetDamage(damage int) *SwordBuilder {
	b.sword.damage = damage
	return b
}

func (b *SwordBuilder) SetLength(length int) *SwordBuilder {
	b.sword.length = length
	return b
}

func (b *SwordBuilder) SetWeight(weight int) *SwordBuilder {
	b.sword.weight = weight
	return b
}

func (b *SwordBuilder) SetPrice(price int) *SwordBuilder {
	b.sword.price = price
	return b
}

func (b *SwordBuilder) Build() *Sword {
	return b.sword
}

func main() {
	swordBuilder := NewSwordBuilder()
	swordBuilder.SetDamage(1000).SetLength(100).
		SetWeight(15).SetPrice(2000)
	newSword := swordBuilder.Build()

	fmt.Printf(" price: %d \n damage: %d \n weight: %d \n lenght: %d \n",
		newSword.price, newSword.damage, newSword.weight, newSword.length)
}
