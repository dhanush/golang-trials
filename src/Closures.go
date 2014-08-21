package main

import (
	"fmt"
)

func makeEven() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEven()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
