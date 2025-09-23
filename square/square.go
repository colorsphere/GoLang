package main

import (
	"fmt"
	"math"
)

type figure interface {
	perimeter() float64
	square() float64
}

type circle struct {
	r float64
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c circle) square() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type triangle struct {
	a float64
	b float64
	c float64
}

func (t triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t triangle) square() float64 {
	p := (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func main() {
	var c figure = circle{1}
	var t figure = triangle{2, 3, 4}

	fmt.Println(t, c)
	fmt.Println(t.perimeter(), c.perimeter())
	fmt.Println(t.square(), c.square())
}
