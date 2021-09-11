package main

import "fmt"

// this is an interface representing a dog that can bark
type Dog interface {
	Bark()
}

type dog struct {
	name string
}

func (d *dog) Bark() {
	fmt.Println(d.name + ": Hup Hup ")
}

func NewDog(name string) Dog {
	return &dog{name}
}

// this is an imposter that we want to be between dogs
// this struct in not accessible for us. so we cannot add a bark method to it
type robot struct {
	name string
}

func NewRobot(name string) *robot {
	return &robot{name}
}

// robot can fabricate itself in the adaptor
type dogAdapter struct {
	robot robot
}

func (d *dogAdapter) Bark() {
	fmt.Println(d.robot.name + ": Hupi Hupi ")
}

func NewDogAdaptor(robot robot) Dog {
	return &dogAdapter{
		robot: robot,
	}
}

func DogsParty(dogs []Dog) {
	for _, d := range dogs {
		d.Bark()
	}
}

func main() {
	shadow := NewDog("shadow")
	lolo := NewDog("lolo")

	r1 := NewRobot("rxjolia625")
	roboDog := NewDogAdaptor(*r1)

	DogsParty([]Dog{shadow, lolo, roboDog})
}
