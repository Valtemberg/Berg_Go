package main

import (
	"fmt"
)

var y string

var z int

func main() {

	y = "Bond, James Bond"

	fmt.Println("Imprimindo valor de y", y, "e pronto.")
	fmt.Printf("%T\n", y)

	fmt.Println("Imprimindo valor de z", z)
	fmt.Printf("%T\n", z)

}
