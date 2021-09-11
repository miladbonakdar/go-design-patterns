package main

import "fmt"

import (
	"bytes"
	"encoding/gob"
)

type CharacterType int

const (
	Agility CharacterType = iota
	Tank
	Healer
)

type Personality struct {
	Armor  int
	Damage int
	Chakra int
}

type Character struct {
	Name        string
	Type        CharacterType
	Personality *Personality
}

func (c *Character) DeepCopy() *Character {
	b := bytes.Buffer{}
	_ = gob.NewEncoder(&b).Encode(c)

	result := Character{}
	_ = gob.NewDecoder(&b).Decode(&result)
	return &result
}

var agilityBase = Character{
	"", Agility, &Personality{50, 100, 20},
}

var healerBase = Character{
	"", Healer, &Personality{40, 20, 100},
}

var tankBase = Character{
	"", Tank, &Personality{100, 40, 40},
}

func NewHero(heroType CharacterType, name string) *Character {
	var base *Character
	switch heroType {
	case Tank:
		base = tankBase.DeepCopy()
		break
	case Agility:
		base = agilityBase.DeepCopy()
		break
	case Healer:
		base = healerBase.DeepCopy()
		break
	default:
		panic("cannot find this character type")
	}
	base.Type = heroType
	base.Name = name
	return base
}

func main() {
	juggernaut := NewHero(Agility, "Juggernaut")
	fmt.Println(juggernaut)

	dazzle := NewHero(Healer, "Dazzle")
	fmt.Println(dazzle)

	warChief := NewHero(Tank, "WarChief")
	fmt.Println(warChief)
}
