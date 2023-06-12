package label

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/juicebottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/waterbottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory"
)

type LabelMachine struct {
	Next bottlefactory.Machine
}

func (m *LabelMachine) Run(b bottle.Bottle) {
	if b.GetLabel() != nil {
		fmt.Println("Label is already added")
		fmt.Println("Bottle is already manufactured")
		return
	}
	label := ""
	switch b.(type) {
	case *waterbottle.WaterBottle:
		label = "paper"
	case *juicebottle.JuiceBottle:
		label = "plastic"
	}
	fmt.Printf("Adding label of type %s\n", label)
	b.SetLabel(&label)
	fmt.Printf("Bottle is ready to use now\n")
}

func (m *LabelMachine) SetNext(machine bottlefactory.Machine) {
	m.Next = machine
}
