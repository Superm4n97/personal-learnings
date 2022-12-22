package main

import "fmt"

type Mammal interface {
	sound() string
	legs() int
}

func basicInfo(m Mammal) {
	fmt.Println("It can", m.sound())
	fmt.Println("Ii has", m.legs(), "legs")
}

type Owner interface {
	getOwnerName() string
}

func identity(o Owner) {
	fmt.Println(o.getOwnerName())
}

type Human struct {
	lgs                 int
	name, voice, gender string
}

func (h Human) getGender() string {
	return h.gender
}
func (h Human) sound() string {
	return h.voice
}
func (h Human) legs() int {
	return h.lgs
}

type Dog struct {
	lg                 int
	noise, oName, name string
}

func (d Dog) sound() string {
	return d.noise
}
func (d Dog) legs() int {
	return d.lg
}
func (d Dog) getOwnerName() string {
	return d.oName
}

type Cat struct {
	oName string
}

func (c Cat) getOwnerName() string {
	return c.oName
}

func main() {
	d := Dog{4, "bark", "Unknown", "Tommy"}
	g := Human{2, "Jorina", "talk", "girl"}
	c := Cat{"Jorina"}

	basicInfo(d)
	basicInfo(g)
	identity(d)
	identity(c)
}
