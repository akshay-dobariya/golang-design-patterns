package function

import "github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"

type UndoFunction struct {
	app application.Application
}

func (sf *UndoFunction) Execute() {
	sf.app.Undo()
}

func (sf *UndoFunction) SetApp(app application.Application) {
	sf.app = app
}
