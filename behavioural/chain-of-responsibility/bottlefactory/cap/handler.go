package cap

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/juicebottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/waterbottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory"
)

type CapMachine struct {
	Next bottlefactory.Machine
}

func (m *CapMachine) Run(b bottle.Bottle) {
	if b.GetCap() != nil {
		fmt.Println("Cap is already added")
		m.Next.Run(b)
		return
	}
	cap := ""
	switch b.(type) {
	case *waterbottle.WaterBottle:
		cap = "plastic"
	case *juicebottle.JuiceBottle:
		cap = "wood"
	}
	fmt.Printf("Adding cap of type %s\n", cap)
	b.SetCap(&cap)
	m.Next.Run(b)
}

func (m *CapMachine) SetNext(machine bottlefactory.Machine) {
	m.Next = machine
}
