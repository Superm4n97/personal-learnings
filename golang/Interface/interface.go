package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type ShapeWithVolume interface {
	Volume() float64
}

type rectangle struct {
	h, w float64
}

func (r rectangle) Area() float64 {
	return r.w * r.h
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.w + r.h)
}

type circle struct {
	r float64
}

func (c circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c circle) PrintRadius() {
	fmt.Println(c.r)
}

type Square struct {
	line float64
}

func (s Square) Area() float64 {
	return s.line * s.line
}

func (s Square) Perimeter() float64 {
	return s.line * 4
}

type Cube struct {
	h, w, d float64
}

func (c Cube) Volume() float64 {
	return 1
}

func (c Cube) Area() float64 {
	return 1
}

func (c Cube) Perimeter() float64 {
	return 1
}

func GetCircleShape(r float64) Shape {
	return circle{r: r}
}

func GetRectangle(w, h float64) Shape {
	return rectangle{
		h: h,
		w: w,
	}
}

func PrintShape(s Shape) {
	fmt.Println("Area : ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}

func PrintShapeVolume(s ShapeWithVolume) {
	fmt.Println("Volume: ", s.Volume())
}

func Print(a interface{}) {
	fmt.Println(reflect.TypeOf(a))
}

func main() {
	c := GetCircleShape(2)

	PrintShape(c)

	r := GetRectangle(2, 4)
	PrintShape(r)

	s := Square{line: 4}
	PrintShape(s)

	cb := Cube{
		h: 1,
		w: 1,
		d: 1,
	}
	PrintShape(cb)

	PrintShapeVolume(cb)

	Print(1)
	Print("name")

}
