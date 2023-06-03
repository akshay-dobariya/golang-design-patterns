package computerbuilder

import "github.com/akshay-dobariya/golang-design-patterns/creational/builder/computer"

type ComputerBuilder interface {
	SetMemoryInMB()
	SetCPUCores()
	SetStorageInGB()
	GetComputer() computer.Computer
}
