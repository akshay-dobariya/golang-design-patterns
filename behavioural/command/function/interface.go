package function

import "github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"

type Function interface {
	Execute()
	SetApp(application.Application)
}
