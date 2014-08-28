package main

import (
	"fmt"
	"os/exec"
)


func main() {
	out, err := exec.Command("free").Output();
	if err!=nil {
		fmt.Println("Error accessing memory info")
	}
	data:=string(out)
	fmt.Println(data)
}
