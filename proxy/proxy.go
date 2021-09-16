package main

import (
	"errors"
	"fmt"
	"time"
)

type MachineController interface {
	PowerOn()
	PowerOff()
	Move(x, y int) error
}

type machineController struct {
	power bool
	x     int
	y     int
}

func (m *machineController) PowerOn() {
	m.power = true
}

func (m *machineController) PowerOff() {
	m.power = false
}

func (m *machineController) Move(x, y int) error {
	if !m.power {
		return errors.New("machine has not turned on")
	}
	m.x = x
	m.y = y
	return nil
}

type enhancedController struct {
	MachineController
	startedAt *time.Time
	stoppedAt *time.Time
}

func (c *enhancedController) PowerOn() {
	// extra behavior has been added to this function
	if c.stoppedAt != nil && time.Since(*c.stoppedAt).Minutes() < 2 {
		fmt.Println("device stopped already. please wait and try again later")
		return
	}
	c.startedAt = TimePointer(time.Now())
	fmt.Println("turning the machine on") // also we can add some logging. have fun we are in the proxy :)
	c.MachineController.PowerOn()
	fmt.Println("machine has turned on")
}

func (c *enhancedController) PowerOff() {
	if c.startedAt != nil && time.Since(*c.startedAt).Minutes() < 2 {
		fmt.Println("device started already. please wait and try again later")
		return
	}
	c.stoppedAt = TimePointer(time.Now())
	fmt.Println("turning the machine off")
	c.MachineController.PowerOff()
	fmt.Println("machine has turned off")
}

func (c *enhancedController) Move(x, y int) error {
	fmt.Printf("trying to move machine to position (%d,%d) \n", x, y)
	return c.MachineController.Move(x, y)
}

func NewController(baseController MachineController) MachineController {
	return &enhancedController{
		MachineController: baseController,
		startedAt:         nil,
		stoppedAt:         nil,
	}
}

func TimePointer(t time.Time) *time.Time {
	return &t
}

func main() {
	// controller proxy
	controller := NewController(&machineController{} /* base controller */)
	controller.PowerOn()
	_ = controller.Move(23, 69)
	controller.PowerOff()
}
