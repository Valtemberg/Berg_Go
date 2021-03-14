// Altere o código a seguir e inclua o seguinte:
// - Crie uma nova variável chamada nota1
// - Crie uma variável chamada nota2
// - Crie uma variável chamada nota3
// - Crie uma variável chamada media:
// -- A variável media deverá receber o valor da media entre nota1, nota2 e nota3
// - Imprima o resultado da media

package main

import (
	"fmt"
)

func main() {
	nota1 := 9
	nota2 := 8
	nota3 := 10
	media := (nota1 + nota2 + nota3) / 3
	soma := nota1 + nota2 + nota3

	fmt.Println("O resultado de soma é:", soma)
	fmt.Println("O valor da média das 3 notas é:", media)
}
