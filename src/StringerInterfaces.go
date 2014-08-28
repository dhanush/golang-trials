package main

import "fmt"
import "sort"

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

type Organs []* Organ

func (s Organs) Len() int { return len(s)}

func (s Organs) Swap(i,j int) { s[i], s[j] = s[j],s[i]}

type ByWeight struct {
	Organs
}

func (s ByWeight) Less (i,j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main () {

	s:=[]*Organ {{"brain", 340}, {"hand", 89}}
	for _,organ:= range s{
		fmt.Println(organ)
	}
	sort.Sort(ByWeight {s})
	fmt.Println(" Sorted " , s)

}
