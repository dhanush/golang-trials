package main

import "fmt"

type Circle struct {
	x,y,r float64
}

func main(){

	var c Circle;
	fmt.Println(c)

	circ:= new (Circle)
	(*circ).x=9
	fmt.Println(circ)

	circ2:= Circle{0,0,8}
	fmt.Println(circ2)
}
