package laptop

import "github.com/akshay-dobariya/golang-design-patterns/creational/builder/computer"

type Laptop struct {
	memoryInMB  int
	cpuCores    int
	storageInGB int
}

func GetLaptop() Laptop {
	return Laptop{}
}

func (l *Laptop) SetMemoryInMB() {
	l.memoryInMB = 4
}

func (l *Laptop) SetCPUCores() {
	l.cpuCores = 2
}

func (l *Laptop) SetStorageInGB() {
	l.storageInGB = 2048
}

func (l *Laptop) GetComputer() computer.Computer {
	computer := computer.Computer{}
	computer.SetData(l.memoryInMB, l.cpuCores, l.storageInGB)
	return computer
}
