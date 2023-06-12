package main

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottle"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory/cap"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory/content"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory/label"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/chain-of-responsibility/bottlefactory/material"
)

func PrintError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

func DescribeBottle(b bottle.Bottle) {
	fmt.Println("==============")
	description := "The bottle is made up of material: '%s', filled with content: '%s', has a cap of type: '%s', and label made up of '%s'\n"
	material := b.GetMaterial()
	content := b.GetContent()
	cap := b.GetCap()
	label := b.GetLabel()
	fmt.Printf(description, *material, *content, *cap, *label)
	fmt.Println("==============")
}

func GetAllMachines() []bottlefactory.Machine {
	// creating all the machines
	materialMachine := &material.MaterialMachine{}
	contentMachine := &content.ContentMachine{}
	capMachine := &cap.CapMachine{}
	labelMachine := &label.LabelMachine{}

	return []bottlefactory.Machine{
		materialMachine, contentMachine, capMachine, labelMachine,
	}
}

func main() {
	machines := GetAllMachines()
	bottleFactory, err := bottlefactory.GetBottleFactory(&machines)
	if err != nil {
		PrintError(err)
	}

	waterBottle, err := bottle.NewBottle("water")
	if err != nil {
		PrintError(err)
	}
	bottleFactory.Run(waterBottle)
	DescribeBottle(waterBottle)

	juiceBottle, err := bottle.NewBottle("juice")
	if err != nil {
		PrintError(err)
	}
	bottleFactory.Run(juiceBottle)
	DescribeBottle(juiceBottle)
}
