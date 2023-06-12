package material

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/juicebottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/waterbottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory"
)

type MaterialMachine struct {
	Next bottlefactory.Machine
}

func (m *MaterialMachine) Run(b bottle.Bottle) {
	if b.GetMaterial() != nil {
		fmt.Println("Material is already added")
		m.Next.Run(b)
		return
	}
	material := ""
	switch b.(type) {
	case *waterbottle.WaterBottle:
		material = "plastic"
	case *juicebottle.JuiceBottle:
		material = "glass"
	}
	fmt.Printf("Adding material of type %s\n", material)
	b.SetMaterial(&material)
	m.Next.Run(b)
}

func (m *MaterialMachine) SetNext(machine bottlefactory.Machine) {
	m.Next = machine
}
