package main

import "fmt"

type Gender bool

const (
	Male   Gender = true
	Female Gender = false
)

type Cat struct {
	CatType string
	Gender  Gender
	Name    string
	Color   string
	Fluffy  bool
}

// this is a factory function
func NewCat(catType, name, color string, gender Gender, fluffy bool) *Cat {
	return &Cat{catType, gender, name, color, fluffy}
}

// you can create specific factory function
func NewDshCat(name, color string, gender Gender) *Cat {
	return &Cat{"DSH", gender, name, color, false}
}

// this is a factory generator
func NewCatFactory(catType string, fluffy bool) func(name, color string, gender Gender) *Cat {
	return func(name, color string, gender Gender) *Cat {
		return &Cat{catType, gender, name, color, fluffy}
	}
}

func main() {
	nabat := NewCat("Persian", "nabat", "orange", Female, true)
	fmt.Println(nabat)

	nero := NewDshCat("nero", "black", Male)
	fmt.Println(nero)

	// first you create a factory from the generator
	persianCatFactory := NewCatFactory("Persian", true)

	asal := persianCatFactory("asal", "orange", Female)
	zoghal := persianCatFactory("zoghal", "black", Male)

	fmt.Println(asal)
	fmt.Println(zoghal)
}
