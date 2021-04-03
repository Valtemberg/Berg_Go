package main

import (
	"fmt"
)

// 1. Crie o seu próprio TIPO. O tipo original ou primário
// deverá ser int.
// 2. Crie uma VARIÁVEL do seu novo TIPO com o IDENTIFICADOR "x"
// utilizando uma declaração normal com "var". (Não é declaração
// curta).
type mengo int

var x mengo

func main() {

	x = 8

	// 3. Na função main:
	// 	a. imprima o valor da variável "x"
	fmt.Println("Valor de x:", x)
	// 	b. imprima o tipo da variável "x"
	fmt.Printf("Tipo de x: %T\n", x)
	// 	c. atribua 42 para a variável "x" utilizando o operador "="
	x = 42
	fmt.Println("Novo valor de x:", x)
	// 	d. imprima o valor da variável "x"

}
