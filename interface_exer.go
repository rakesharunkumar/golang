package main

import (
	"fmt"
)

type square struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	sq := square{
		4,
		4,
	}
	cr := circle{
		2.4,
	}
	fmt.Println(sq, cr)
	//create method to calculate area
	areaofsquare := sq.area()
	fmt.Println(areaofsquare)
	areaofcircle := cr.area()
	fmt.Println(areaofcircle)
	info(sq)
	info(cr)
}

func (s square) area() float64 {
	area := s.length * s.breadth
	return area
	//fmt.Println("Area of Square is ", area)
}
func (c circle) area() float64 {
	area := c.radius * 3.14 * c.radius
	return area
	//fmt.Println("Area of circle is", area)
}

func info(h shape) {
	fmt.Println(h.area())
}
