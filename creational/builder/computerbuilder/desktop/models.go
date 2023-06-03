package desktop

import "github.com/akshay-dobariya/golang-design-patterns/creational/builder/computer"

type Desktop struct {
	memoryInMB  int
	cpuCores    int
	storageInGB int
}

func GetDesktop() Desktop {
	return Desktop{}
}

func (d *Desktop) SetMemoryInMB() {
	d.memoryInMB = 64
}

func (d *Desktop) SetCPUCores() {
	d.cpuCores = 8
}

func (d *Desktop) SetStorageInGB() {
	d.storageInGB = 102400
}

func (d *Desktop) GetComputer() computer.Computer {
	computer := computer.Computer{}
	computer.SetData(d.memoryInMB, d.cpuCores, d.storageInGB)
	return computer
}
