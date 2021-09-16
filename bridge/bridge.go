package main

import (
	"fmt"
	"log"
)

type Computer interface {
	Print(message string)
	DoOtherStuff()
}

type Printer interface {
	printFile(message string)
}

type mac struct {
	printer Printer
	name    string
}

func (m *mac) Print(message string) {
	if m.printer == nil {
		log.Fatal("computer " + m.name + " doesn't have a printer")
	}
	println("mac printing")
	m.printer.printFile(message)
}

func (m *mac) DoOtherStuff() {
	println("do mac things")
}

func NewMac(p Printer) Computer {
	return &mac{
		printer: p,
		name:    "Mac",
	}
}

type windows struct {
	printer Printer
	name    string
}

func (m *windows) Print(message string) {
	if m.printer == nil {
		log.Fatal("computer " + m.name + " doesn't have a printer")
	}
	println("windows printing")
	m.printer.printFile(message)
}

func (m *windows) DoOtherStuff() {
	println("do windows things")
}

func NewWindows(p Printer) Computer {
	return &windows{
		printer: p,
		name:    "Windows",
	}
}

type hp struct{}

func (p *hp) printFile(message string) {
	fmt.Println("Printing by a HP Printer")
	fmt.Println(message)
	fmt.Println("---------------------------")
}

type epson struct{}

func (p *epson) printFile(message string) {
	fmt.Println("Printing by a EPSON Printer")
	fmt.Println(message)
	fmt.Println("---------------------------")
}

func main() {
	epson := &epson{}
	hp := &hp{}

	mac := NewMac(epson)        // bridge design pattern flexibly let you create multiple type of computers
	windowsHp := NewWindows(hp) // with different printers
	windowsEpson := NewWindows(epson)

	mac.Print("To live is to risk it all; otherwise you're just an inert chunk of randomly assembled molecules drifting wherever the universe blows you.")
	windowsHp.Print("Nobody exists on purpose. Nobody belongs anywhere. We’re all going to die. Come watch TV.")
	windowsEpson.Print("Sometimes science is more art than science.")
}
