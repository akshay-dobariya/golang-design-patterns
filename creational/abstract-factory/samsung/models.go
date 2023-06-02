package samsung

import (
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/charger"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/mobile"
)

type Samsung struct {
}

type SamsungMobile struct {
	mobile.Mobile
}

type SamsungCharger struct {
	charger.Charger
}

func (m *Samsung) MakeMobile() mobile.MobileInterface {
	samsungMobile := &SamsungMobile{}
	samsungMobile.SetHardware("samsung hardware")
	samsungMobile.SetSoftware("samsung software")
	return samsungMobile
}

func (m *Samsung) MakeCharger() charger.ChargerInterface {
	samsungCharger := &SamsungCharger{}
	samsungCharger.SetWatt(75)
	return samsungCharger
}
