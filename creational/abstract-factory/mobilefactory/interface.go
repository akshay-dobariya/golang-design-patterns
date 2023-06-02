package mobilefactory

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/apple"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/charger"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/mobile"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/samsung"
)

type MobileFactory interface {
	MakeMobile() mobile.MobileInterface
	MakeCharger() charger.ChargerInterface
}

func GetMobileFactory(companyName string) (MobileFactory, error) {
	if companyName == "samsung" {
		return &samsung.Samsung{}, nil
	} else if companyName == "apple" {
		return &apple.Apple{}, nil
	}
	return nil, fmt.Errorf("Couldn't find factory named %s", companyName)
}
