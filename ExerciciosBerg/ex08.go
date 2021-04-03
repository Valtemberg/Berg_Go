package main

import (
	"fmt"
)

// Utilizando o código do exemplo anterior:
// 1. no escopo de pacote, utilizando a palavra chave "var",
// crie uma VARIÁVEL com o IDENTIFICADOR "y". A variável
// deverá ser do tipo original ou que dá origem ao SEU TIPO
// criado anteriormente.
//
type mengo int

var x mengo

func main() {

	x = 8

	var y int

	// 2. na função main:
	// 	a. isto já deverá estar feito:
	// 		i. imprimir o valor de "x"
	// 		ii. imprimir o tipo de "x"
	// 		iii. atribuir o VALOR à variável "x" utilizando "="
	// 		iv. imprimir o valor de "x"
	fmt.Println("Valor de x:", x)
	fmt.Printf("Tipo de x: %T\n", x)
	x = 42
	fmt.Println("Novo valor de x:", x)

	// 	b. Agora faça:
	// 		i. utilize CONVERSÃO para converter o TIPO do valor
	// 		de "x" para o tipo original de seu TIPO.
	// 		ii. Atribuir este valor a "y".
	y = int(x)
	// 		iii. Imprima o tipo de "y".
	fmt.Printf("Tipo de y: %T\n", y)

}
