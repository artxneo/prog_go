package main

import "fmt"

type Animal interface {
	foodNeedle
	fmt.Stringer
	animalWeight
}

type foodNeedle interface {
	monthFoodsNeed() float64
}

type animalWeight interface {
	getWeight() float64
}

type dog struct {
	weight float64
	name   string
}

type cat struct {
	weight float64
	name   string
}

type cow struct {
	weight float64
	name   string
}

func (d dog) monthFoodsNeed() float64 {
	const dogFoodsNeed float64 = 10
	const dogFoodsNeedForWeight float64 = 5
	dogFoodCoefficient := dogFoodsNeed / dogFoodsNeedForWeight

	return d.weight * dogFoodCoefficient
}

func (c cat) monthFoodsNeed() float64 {
	const catFoodCoefficient float64 = 7

	return c.weight * catFoodCoefficient
}

func (c cow) monthFoodsNeed() float64 {
	const cowFoodCoefficient float64 = 25

	return c.weight * cowFoodCoefficient
}

func (d dog) String() string {

	return d.name
}

func (c cat) String() string {
	return c.name
}

func (c cow) String() string {
	return c.name
}

func (d dog) getWeight() float64 {
	return d.weight
}

func (c cat) getWeight() float64 {
	return c.weight
}

func (c cow) getWeight() float64 {
	return c.weight
}
