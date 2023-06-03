package maruti

import (
	"github.com/akshay-dobariya/golang-design-patterns/creational/factory/car"
)

type MarutiCar struct {
	car.Car
}

func NewMarutiCar() car.Car {
	car := car.Car{}
	car.SetCarBodyMaterial("plastic")
	car.SetCarWheels("normal")
	return car
}
