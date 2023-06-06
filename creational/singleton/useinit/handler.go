package useinit

import "fmt"

type Singleton struct {
}

var singletonInstance *Singleton

func init() {
	fmt.Println("useinit - Creating Singleton Instance Now")
}

func GetSingletonInstance() *Singleton {
	return singletonInstance
}
