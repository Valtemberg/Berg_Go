package main

import (
	"fmt"
)

var y = 54

func main() {
	fmt.Println("Olá mundo")
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Printf("%b\n", y)  //em binarios
	fmt.Printf("%x\n", y)  //em hexadecimal
	fmt.Printf("%#x\n", y) //hexadecimal com 0x na frente
	fmt.Printf("%X\n", y)  //hexadecimal em maiusculas

	x := "Berg torce Flamengo"

	fmt.Printf("%#x\n%b\n%x\n", x, x, x)

}
