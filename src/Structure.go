package main

import "fmt"
import "math"

type Circle struct {
	x,y,r float64
}

func area(circle *Circle) float64 {

	return math.Pi * circle.r * circle.r
}

func (c *Circle) rArea () float64 {
	return math.Pi * c.r * c.r
}

func main(){

	var c Circle;
	fmt.Println(c)

	circ:= new (Circle)
	(*circ).x=9
	(*circ).r=1
	fmt.Println(circ)

	circ2:= Circle{0,0,8}
	fmt.Println(circ2)

	a:= area(&circ2);
	fmt.Println("Area is ", a)

	b:= area(circ)
	fmt.Println(" Are of b is ", b)
	
	fmt.Println("Area from method ", circ.rArea())


}
