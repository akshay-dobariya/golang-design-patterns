package main

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/factory/carfactory"
)

func PrintError(err error) {
	fmt.Printf("Error: %s", err.Error())
}

func main() {
	marutiCar, err := carfactory.GetCar("maruti")
	if err != nil {
		PrintError(err)
	}
	marutiCar.PrintCarDetails()

	tataCar, err := carfactory.GetCar("tata")
	if err != nil {
		PrintError(err)
	}
	tataCar.PrintCarDetails()
}
