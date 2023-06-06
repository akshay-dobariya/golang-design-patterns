package useonce

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var once sync.Once

var singletonInstance *Singleton

func GetSingletonInstance() *Singleton {
	if singletonInstance == nil {
		once.Do(func() {
			fmt.Println("useonce - Creating Singleton Instance Now")
			singletonInstance = &Singleton{}
		})
	} else {
		fmt.Println("useonce - Singleton Instance is already created")
	}
	return singletonInstance
}
