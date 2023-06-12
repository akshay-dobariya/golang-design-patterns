package bottle

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/juicebottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle/waterbottle"
)

type Bottle interface {
	SetMaterial(*string)
	SetContent(*string)
	SetCap(*string)
	SetLabel(*string)
	GetMaterial() *string
	GetContent() *string
	GetCap() *string
	GetLabel() *string
}

func NewBottle(bottleType string) (Bottle, error) {
	if bottleType == "juice" {
		return &juicebottle.JuiceBottle{}, nil
	} else if bottleType == "water" {
		return &waterbottle.WaterBottle{}, nil
	} else {
		return nil, fmt.Errorf("We don't manufacture %s type of bottles", bottleType)
	}
}
