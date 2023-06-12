package content

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/juicebottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/waterbottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory"
)

type ContentMachine struct {
	Next bottlefactory.Machine
}

func (m *ContentMachine) Run(b bottle.Bottle) {
	if b.GetContent() != nil {
		fmt.Println("Content is already added")
		m.Next.Run(b)
		return
	}
	content := ""
	switch b.(type) {
	case *waterbottle.WaterBottle:
		content = "water"
	case *juicebottle.JuiceBottle:
		content = "juice"
	}
	fmt.Printf("Adding content of type %s\n", content)
	b.SetContent(&content)
	m.Next.Run(b)
}

func (m *ContentMachine) SetNext(machine bottlefactory.Machine) {
	m.Next = machine
}
