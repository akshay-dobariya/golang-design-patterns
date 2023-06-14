package collection

import (
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/iterator"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/objects"
)

type MobileCollection struct {
	Mobiles *[]objects.Mobile
}

func (uc *MobileCollection) CreateInterator() iterator.Interator {
	mobileIterator := &objects.MobileIterator{
		Mobiles: uc.Mobiles,
	}
	return mobileIterator
}
