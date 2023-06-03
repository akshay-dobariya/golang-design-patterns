package computer

import "fmt"

type Computer struct {
	memoryInMB  int
	cpuCores    int
	storageInGB int
}

func (c *Computer) SetData(memory, cores, storage int) {
	c.memoryInMB = memory
	c.cpuCores = cores
	c.storageInGB = storage
}

func (c *Computer) PrintData() {
	fmt.Println("=============")
	fmt.Println("Memory in MB: ", c.memoryInMB)
	fmt.Println("CPU cores: ", c.cpuCores)
	fmt.Println("Storage in GB: ", c.storageInGB)
	fmt.Println("=============")
}
