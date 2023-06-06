package main

import (
	"fmt"
	"sync"

	"github.com/akshay-dobariya/golang-design-patterns/creational/singleton/useinit"
	"github.com/akshay-dobariya/golang-design-patterns/creational/singleton/usemutex"
	"github.com/akshay-dobariya/golang-design-patterns/creational/singleton/useonce"
)

func main() {
	wg := sync.WaitGroup{}

	fmt.Println("================")
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			usemutex.GetSingletonInstance()
		}(&wg)
	}
	wg.Wait()
	fmt.Println("================")
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			useinit.GetSingletonInstance()
		}(&wg)
	}
	wg.Wait()
	fmt.Println("================")
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			useonce.GetSingletonInstance()
		}(&wg)
	}
	wg.Wait()
	fmt.Println("================")
}
