package main

import (
	"fmt"
)

func main() {
	// 1. Utilizando o operador para declaração curta (ou short) de variável,
	// ATRIBUA os seguintes VALORES para os IDENTIFICADORES "x", "y" and "z":

	// a. 42
	// b. "James Bond"
	// c. true

	x := 42
	y := "James Bond"
	z := true

	// 2. Agora imprima os valores armazenados nestas variáveis
	// utilizando:
	// a. Uma única chamada de Print (Um único Print)
	fmt.Println("Valor de x = ", x, ", valor de y = ", y, "e valor de z =", z)

	// b. Múltiplas chamadas de Print
	fmt.Println("Valor de x = ", x)
	fmt.Println("Valor de y = ", y)
	fmt.Println("Valor de z = ", z)

}
