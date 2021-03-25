package main

import (
	"fmt"
)

//Declaração de variaveis:
//variaveis com declaração curta: :=
//Variavel com declaração completa: =

func main(){
	x := 42 //Declaramos variavel e atribuimos valor
	fmt.Println(x)

	x = 99 //Atribuimos novo valor
	fmt.Println(x)

	y := 100 * 24 / 2 //Declaramos variavel e atribuimos valor da expressão
	fmt.Println(y)

	z := "Bond, James Bond"
	codigo := "007"

	fmt.Println("Meu nome é", z, "mas você pode me chamar de:", codigo)
}