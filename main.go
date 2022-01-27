package main

import (
	"fmt"

	"example.com/prueba1/selfFuncs"
)

func main() {

	var a int = 1

	b := &a
	fmt.Println(*b)
	*b = 100
	fmt.Println(a, b)

	myPC := selfFuncs.Pc{Ram: 12, Disk: 200, Brand: "asus"}
	laptop := selfFuncs.Pc{Ram: 8, Disk: 100, Brand: "dell"}
	fmt.Println(myPC)
	myPC.Ping()
	myPC.DuplicateRAM()
	fmt.Println(myPC)
	laptop.DuplicateRAM()
	fmt.Println(laptop)
}
