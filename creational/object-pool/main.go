package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/akshay-dobariya/golang-design-patterns/creational/object-pool/pool"
	"github.com/akshay-dobariya/golang-design-patterns/creational/object-pool/poolobject"
	"github.com/akshay-dobariya/golang-design-patterns/creational/object-pool/poolobject/connection"
)

func PrintError(i int, err error) {
	fmt.Printf("Error in Thead %d: %s\n", i, err.Error())
}

func ConsumeConnections(i int, p *pool.Pool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		myObj, err := p.GetPoolObject()
		if err != nil {
			PrintError(i, err)
		} else {
			fmt.Printf("Thread: %d - pool object id: %d\n", i, (*myObj).GetId())
			time.Sleep(1)
			p.FreePoolObject(myObj)
			return
		}
	}

}

func main() {
	connections := []poolobject.PoolObjectInterface{}
	for i := 0; i < 3; i++ {
		c := connection.Connection{}
		c.SetId(i)
		connections = append(connections, &c)
	}
	p, _ := pool.InitPool(connections)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go ConsumeConnections(i, p, &wg)
	}
	wg.Wait()
}
