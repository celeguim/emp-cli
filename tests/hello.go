package main

import (
	"fmt"
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func TestMain(t *testing.T) {
	fmt.Println("hello")
	t.Log("OK")
}

func calculate(s Shape) float64 {
	return s.Area()
}

func main() {
	fmt.Println("Area calculator v0.1")

	r := Rectangle{width: 4, height: 4}
	fmt.Printf("Rectangle area: %f \n", calculate(r))

	c := Circle{radius: 3}
	fmt.Printf("Circle area: %f \n", calculate(c))
}
