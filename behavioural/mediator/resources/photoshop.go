package resources

import (
	"fmt"
	"time"
)

type Photoshop struct {
	name           string
	memoryRequired MemorySizeInMB
	timeRequired   int
	resource       Resource
}

func (p *Photoshop) MemoryRequired() MemorySizeInMB {
	return p.memoryRequired
}

func (p *Photoshop) RequestMemory() {
	if !p.resource.IsAvailable(p) {
		fmt.Printf("Waiting for Memory: %s\n", p.name)
	}
}

func (p *Photoshop) UseAndFreeMemory() {
	go func(second int) {
		time.Sleep(time.Duration(second * int(time.Second)))
		p.resource.NotifyFree(p)
	}(p.timeRequired)
}

func (p *Photoshop) GotMemory() {
	fmt.Printf("Got Memory: %s\n", p.name)
	p.UseAndFreeMemory()
}

func (p *Photoshop) GetAppName() string {
	return p.name
}
