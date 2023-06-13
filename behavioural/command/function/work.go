package function

import (
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"
)

type WorkFunction struct {
	inProgress string
	app        application.Application
}

func (sf *WorkFunction) Execute() {
	sf.app.Work(sf.inProgress)
}

func (sf *WorkFunction) AddWork(work string) {
	sf.inProgress = work
}

func (sf *WorkFunction) SetApp(app application.Application) {
	sf.app = app
}
