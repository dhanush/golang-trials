package main

import (
	"fmt"
	"time"
)

func print10(name string, n int) {
	fmt.Println("Printing for ", name)
	for i:=0; i<n ;i++ {
		fmt.Println(i)
	}
}

func main() {
	go print10("Dhanush", 9)
	go print10("Gopinath", 10)
	time.Sleep(time.Second *5)
}
