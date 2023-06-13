package ui

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/command/function"
)

type ActionButton struct {
	function function.Function
}

func (b *ActionButton) Press() {
	b.function.Execute()
}

func GetActionButton(action string) (Button, error) {
	switch action {
	case "save":
		function := function.SaveFunction{}
		function.SetApp(application.GetPhotoshopApp())
		button := ActionButton{
			function: &function,
		}
		return &button, nil
	case "undo":
		function := function.UndoFunction{}
		function.SetApp(application.GetPhotoshopApp())
		button := ActionButton{
			function: &function,
		}
		return &button, nil
	case "display":
		function := function.DisplayWorkFunction{}
		function.SetApp(application.GetPhotoshopApp())
		button := ActionButton{
			function: &function,
		}
		return &button, nil
	default:
		return nil, fmt.Errorf("Action button is not available for action type: %s", action)
	}
}
