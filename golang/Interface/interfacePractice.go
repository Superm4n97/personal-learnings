package main

import "fmt"

type vehicle interface {
	speed() int
}

type twoWheelers interface {
	seat() int
	headlight() int
}

type fourWheelers interface {
	doors() int
	model() string
}

type Car struct {
	sp                      int
	company, signatureModel string
}

func (t Car) speed() int {
	return t.sp
}
func (t Car) doors() int {
	return 4
}
func (t Car) model() string {
	return t.company + t.signatureModel
}

type biCycle struct {
	sp, st, tierSize int
}

func (b biCycle) speed() int {
	return b.sp
}
func (b biCycle) seat() int {
	return b.st
}
func (b biCycle) headlight() int {
	return 0
}

type bike struct {
	sp, st int
	mdl    string
}

func (b bike) speed() int {
	return b.sp
}
func (b bike) seat() int {
	return b.st
}
func (b bike) headlight() int {
	return 1
}

func basicInfoFourWheel(f fourWheelers) {
	fmt.Println("Basic Info:")
	f.model()
	f.doors()
	fmt.Println("high price")
}

func basicInfoTwoWheel() {

}
