package main

import (
	"fmt"
	"time"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/mediator/resources"
)

func main() {
	applications := make(chan resources.Application)
	for i := 1; i <= 5; i++ {
		myAppCreator := func(appType string, i int) {
			app := resources.GetApplication(appType, fmt.Sprintf("%s:%d", appType, i), 1024*((i%3)+1), i+2)
			applications <- app
		}
		go myAppCreator("photoshop", i)
		go myAppCreator("illustrator", i)
	}

	for i := 0; i < 10; i++ {
		app := <-applications
		go app.RequestMemory()
	}

	time.Sleep(100 * time.Second)
}
