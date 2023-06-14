package collection

import (
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/iterator"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/objects"
)

type UserCollection struct {
	Users *[]objects.User
}

func (uc *UserCollection) CreateInterator() iterator.Interator {
	userIterator := &objects.UserIterator{
		Users: uc.Users,
	}
	return userIterator
}
