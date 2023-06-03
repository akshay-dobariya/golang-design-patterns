package carfactory

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/factory/car"
	"github.com/akshay-dobariya/golang-design-patterns/creational/factory/car/maruti"
	"github.com/akshay-dobariya/golang-design-patterns/creational/factory/car/tata"
)

func GetCar(carType string) (car.CarInterface, error) {
	if carType == "maruti" {
		car := maruti.NewMarutiCar()
		return &car, nil
	} else if carType == "tata" {
		car := tata.NewTataCar()
		return &car, nil
	}
	return nil, fmt.Errorf("Car type: %s is not valid", carType)
}
