package usemutex

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var mutex sync.Mutex

var singletonInstance *Singleton

func GetSingletonInstance() *Singleton {
	if singletonInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singletonInstance == nil {
			fmt.Println("usemutex - Creating Singleton Instance Now")
			singletonInstance = &Singleton{}
		} else {
			fmt.Println("usemutex - Singleton Instance is now created some other goroutine")
		}
	} else {
		fmt.Println("usemutex - Singleton Instance is already created")
	}
	return singletonInstance
}
