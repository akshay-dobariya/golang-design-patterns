package main

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/charger"
	"github.com/akshay-dobariya/golang-design-patterns/creational/abstract-factory/mobilefactory/mobile"
)

func PrintError(err error) {
	fmt.Printf("Error: %s", err.Error())
}

func PrintDeviceSet(mobile mobile.MobileInterface, charger charger.ChargerInterface) {
	fmt.Println("==============")
	fmt.Printf("Mobile: Hardware: %s & Software: %s\n", mobile.GetHardware(), mobile.GetSoftware())
	fmt.Printf("Charger: Watt: %d\n", charger.GetWatt())
	fmt.Println("==============")
}

func main() {
	samsungFactory, err := mobilefactory.GetMobileFactory("samsung")
	if err != nil {
		PrintError(err)
	}
	samsungMobile := samsungFactory.MakeMobile()
	samsungCharger := samsungFactory.MakeCharger()

	PrintDeviceSet(samsungMobile, samsungCharger)

	appleFactory, err := mobilefactory.GetMobileFactory("apple")
	if err != nil {
		PrintError(err)
	}
	appleMobile := appleFactory.MakeMobile()
	appleCharger := appleFactory.MakeCharger()
	PrintDeviceSet(appleMobile, appleCharger)
}
