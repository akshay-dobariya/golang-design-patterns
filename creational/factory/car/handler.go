package car

import "fmt"

type Car struct {
	wheelsType   string
	bodyMaterial string
}

func (tc *Car) SetCarBodyMaterial(material string) {
	tc.bodyMaterial = material
}

func (tc *Car) SetCarWheels(wheelType string) {
	tc.wheelsType = wheelType
}

func (tc *Car) PrintCarDetails() {
	fmt.Println("============")
	fmt.Println("Car WheelType:", tc.wheelsType)
	fmt.Println("Car BodyMaterial:", tc.bodyMaterial)
	fmt.Println("============")
}
