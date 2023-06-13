package main

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/command/ui"
)

func PrintError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

func main() {
	saveButton, err := ui.GetActionButton("save")
	if err != nil {
		PrintError(err)
	}
	undoButton, err := ui.GetActionButton("undo")
	if err != nil {
		PrintError(err)
	}
	displayButton, err := ui.GetActionButton("display")
	if err != nil {
		PrintError(err)
	}
	workButton, err := ui.GetWorkButton("work")
	if err != nil {
		PrintError(err)
	}

	fmt.Println("=========== checking empty data ===========")
	saveButton.Press()
	undoButton.Press()
	displayButton.Press()
	fmt.Println("=========== checking work ===========")
	workButton.(*ui.WorkButton).AddWork("photoshop")
	workButton.(*ui.WorkButton).Press()
	workButton.(*ui.WorkButton).AddWork("is")
	workButton.(*ui.WorkButton).Press()
	workButton.(*ui.WorkButton).AddWork("working")
	workButton.(*ui.WorkButton).Press()
	workButton.(*ui.WorkButton).AddWork("fine")
	workButton.(*ui.WorkButton).Press()
	workButton.(*ui.WorkButton).AddWork("!")
	workButton.(*ui.WorkButton).Press()
	displayButton.Press()
	undoButton.Press()
	displayButton.Press()
	saveButton.Press()
	displayButton.Press()
}
