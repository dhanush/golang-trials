package main

import (
	"fmt"
)

func smallest(x []int) int {
	small := x[0]
	for _, value := range x {
		if value < small {
			small = value
		}
	}
	return small
}

func main() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	small := smallest(x)

	fmt.Println("Smallest Number is ", small)

}
