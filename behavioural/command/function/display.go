package function

import "github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"

type DisplayWorkFunction struct {
	app application.Application
}

func (sf *DisplayWorkFunction) Execute() {
	sf.app.DisplayWork()
}

func (sf *DisplayWorkFunction) SetApp(app application.Application) {
	sf.app = app
}
