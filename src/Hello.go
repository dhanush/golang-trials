package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	x := make([]float64, 5)
	fmt.Println(x)
	x[1] = 123

	fmt.Println(x)

	//	var mapper map[string] int;

	mapper := make(map[string]int)

	mapper["key1"] = 190
	mapper["key2"] = 191

	fmt.Println(mapper)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	fmt.Println(elements["Un"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)
	
	if name, ok:= elements["F"]; ok {
		fmt.Println(name, ok)
	}
}
