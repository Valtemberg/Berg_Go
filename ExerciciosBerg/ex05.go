package main

import (
	"fmt"
)

// Utilizando o código do exercício 04...
// 1. No escopo de pacote, atribua os seguintes valores
// para as 3 variáveis:
//
// a. para x, atribua 42
// b. para y, atribua "James Bond"
// c. para z, atribua true

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	// 2. Na função main:
	//
	// a. utilize fmt.Sprintf para imprimir todos os VALORES em
	// uma única string.ATRIBUA o valor retornado do TIPO string
	// utilizando o declaradaor curto (short) para uma VARIÁVEL
	// com o IDENTIFICADOR "s".

	s := fmt.Sprintf("%x\t%y\t%z\n", x, y, z)

	// b. imprima o valor de "s"
	fmt.Println(s)

}
