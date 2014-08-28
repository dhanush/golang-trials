package main

import "fmt"

type Organ struct {
	Name string
	Weight Grams
}

type Grams int

func (o *Organ) String() string {
	return fmt.Sprintf("%v (%v)", o.Name, o.Weight)
}

func (g Grams) String() string {
	return fmt.Sprintf("%dg", int(g))
}

func main () {

	s:=[]*Organ {{"brain", 340}, {"hand", 89}}
	for _,organ:= range s{
		fmt.Println(organ)
	}
}
