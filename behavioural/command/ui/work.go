package ui

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/command/application"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/command/function"
)

type WorkButton struct {
	function function.Function
}

func (b *WorkButton) Press() {
	b.function.Execute()
}

func (b *WorkButton) AddWork(work string) {
	b.function.(*function.WorkFunction).AddWork(work)
}

func GetWorkButton(action string) (Button, error) {
	if action == "work" {
		function := function.WorkFunction{}
		function.SetApp(application.GetPhotoshopApp())
		button := WorkButton{
			function: &function,
		}
		return &button, nil
	}
	return nil, fmt.Errorf("Work button is not available for action type: %s", action)
}
