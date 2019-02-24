package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {

	return s.length * s.length

}
func (c *circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area is", s.area())
}



func main() {



	c1 := circle{
		radius: 12.0,
	}
	fmt.Printf("%T\n", &c1)

	info(&c1)
}

