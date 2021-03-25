package main

import (
	"fmt"
)

var y int = 43

var z int
var w string
var a bool
var b float64

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	foo()
	fmt.Println("Encerrei foo...")
}

func foo() {
	fmt.Println("Novamente:", y)
	fmt.Println("Valor de z(int) é:", z)
	fmt.Println("valor de w(string) é:", w)
	fmt.Println("Valor de a(bool) é:", a)
	fmt.Println("Valor de b(Float64) é:", b)
}
