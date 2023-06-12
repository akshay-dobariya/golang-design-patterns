package bottlefactory

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle"
)

type Machine interface {
	Run(bottle.Bottle)
	SetNext(Machine)
}

func GetBottleFactory(machines *[]Machine) (Machine, error) {
	if len(*machines) == 0 {
		return nil, fmt.Errorf("Atleast one machine is required to setup a factory")
	}
	for i, machine := range *machines {
		if i == len(*machines)-1 {
			machine.SetNext(nil)
		} else {
			machine.SetNext((*machines)[i+1])
		}
	}
	return (*machines)[0], nil
}
