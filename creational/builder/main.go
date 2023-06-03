package main

import (
	"github.com/akshay-dobariya/golang-design-patterns/creational/builder/computerbuilder/desktop"
	"github.com/akshay-dobariya/golang-design-patterns/creational/builder/computerbuilder/laptop"
	"github.com/akshay-dobariya/golang-design-patterns/creational/builder/director"
)

func main() {
	director := director.Director{}
	laptopBuilder := laptop.GetLaptop()
	director.SetBuilder(&laptopBuilder)
	l := director.ProduceComputer()
	l.PrintData()

	desktopBuilder := desktop.GetDesktop()
	director.SetBuilder(&desktopBuilder)
	d := director.ProduceComputer()
	d.PrintData()
}
