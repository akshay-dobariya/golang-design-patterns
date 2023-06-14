package resources

import (
	"fmt"
	"time"
)

type Illustrator struct {
	name           string
	memoryRequired MemorySizeInMB
	timeRequired   int
	resource       Resource
}

func (i *Illustrator) MemoryRequired() MemorySizeInMB {
	return i.memoryRequired
}

func (i *Illustrator) RequestMemory() {
	if !i.resource.IsAvailable(i) {
		fmt.Printf("Waiting for Memory: %s\n", i.name)
	}
}

func (i *Illustrator) UseAndFreeMemory() {
	go func(second int) {
		time.Sleep(time.Duration(second * int(time.Second)))
		i.resource.NotifyFree(i)
	}(i.timeRequired)
}

func (i *Illustrator) GotMemory() {
	fmt.Printf("Got Memory: %s\n", i.name)
	i.UseAndFreeMemory()
}

func (i *Illustrator) GetAppName() string {
	return i.name
}
