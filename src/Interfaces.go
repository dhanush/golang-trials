package main

import "fmt" 
import "math" 


type Shape interface {
	area() float64 
}

type Circle struct {
	x,y, r float64
}

type Square struct {
	a float64
}

func (s *Square) area() float64 {
	return s.a*s.a 
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}


func main() {

	c:=Circle {0,0,8}
	fmt.Println("Area of Circle ", c.area())

	s:=Square{10}
	fmt.Println("Area of Square ", s.area())
}


