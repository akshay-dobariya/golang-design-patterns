package tata

import (
	"github.com/akshay-dobariya/golang-design-patterns/creational/factory/car"
)

type TataCar struct {
	car.Car
}

func NewTataCar() car.Car {
	car := car.Car{}
	car.SetCarBodyMaterial("steel")
	car.SetCarWheels("alloy")
	return car
}
