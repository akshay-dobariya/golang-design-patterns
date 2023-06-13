package function

import "github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"

type SaveFunction struct {
	app application.Application
}

func (sf *SaveFunction) Execute() {
	sf.app.Save()
}

func (sf *SaveFunction) SetApp(app application.Application) {
	sf.app = app
}
