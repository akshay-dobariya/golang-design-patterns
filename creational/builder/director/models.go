package director

import (
	"github.com/akshay-dobariya/golang-design-patterns/creational/builder/computer"
	"github.com/akshay-dobariya/golang-design-patterns/creational/builder/computerbuilder"
)

type Director struct {
	builder computerbuilder.ComputerBuilder
}

func (d *Director) SetBuilder(builder computerbuilder.ComputerBuilder) {
	d.builder = builder
}

func (d *Director) ProduceComputer() computer.Computer {
	d.builder.SetMemoryInMB()
	d.builder.SetStorageInGB()
	d.builder.SetCPUCores()
	return d.builder.GetComputer()
}
