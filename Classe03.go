package main

import (
	"fmt"
)

//Escopo geral: Declarada fora da função, declarando no pacote
//Escopo de função: Declarando dentro da função

//Variavel curta só sera usado dentro da função
//variavel completa pode ser usado dentro ou fora da função

var z int

func main() {
	x := 42 // declarador curto, declarando variavel x com valor 42
	fmt.Println(x)

	var y = 43
	fmt.Println(y)
	z = 15
	fmt.Println(z)

}
