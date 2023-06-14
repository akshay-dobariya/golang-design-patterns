package collection

import "github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/iterator"

type Collection interface {
	CreateInterator() iterator.Interator
}
