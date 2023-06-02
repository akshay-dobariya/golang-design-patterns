package apple

import (
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/charger"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/mobile"
)

type Apple struct {
}

type AppleMobile struct {
	mobile.Mobile
}

type AppleCharger struct {
	charger.Charger
}

func (m *Apple) MakeMobile() mobile.MobileInterface {
	appleMobile := &AppleMobile{}
	appleMobile.SetHardware("samsung hardware")
	appleMobile.SetSoftware("samsung software")
	return appleMobile
}

func (m *Apple) MakeCharger() charger.ChargerInterface {
	appleCharger := &AppleCharger{}
	appleCharger.SetWatt(12)
	return appleCharger
}
