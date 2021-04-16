package main

import (
	"fmt"
	"math"
)

// Você possui uma turma com 4 alunos cujos nomes são Pedro, Ricardo, Paula e Marcia.
// Cada um fez 3 exames e tiveram notas diversas, porém a média final de cada deverá ser igual ou maior a 7.0 para que sejam aprovados.
// As notas de cada um deles são as seguintes:
// - Pedro: 4.5, 7.9 e 8.1
// - Ricardo: 5.7, 6.0 e 9.1
// - Paula: 6.8, 7.6 e 8.7
// - Marcia: 7.5, 5.9 e 7.6

// a. Crie uma variável que receba o nome de cada aluno. Estas variáveis devem possuir escopo geral.

var aluno01 string = "Pedro"
var aluno02 string = "Ricardo"
var aluno03 string = "Paula"
var aluno04 string = "Marcia"
var media01 float64
var media02 float64
var media03 float64
var media04 float64

func main() {

	// b. Para cada nota de cada aluno, crie uma variável. Estas variáveis de nota precisam ser com escopo de função.

	aluno01n1 := 4.5
	aluno01n2 := 7.9
	aluno01n3 := 8.1

	aluno02n1 := 5.7
	aluno02n2 := 6.0
	aluno02n3 := 9.1

	aluno03n1 := 6.8
	aluno03n2 := 7.6
	aluno03n3 := 8.7

	aluno04n1 := 7.5
	aluno04n2 := 5.9
	aluno04n3 := 7.6

	media01 = (aluno01n1 + aluno01n2 + aluno01n3) / 3
	media02 = (aluno02n1 + aluno02n2 + aluno02n3) / 3
	media03 = (aluno03n1 + aluno03n2 + aluno03n3) / 3
	media04 = (aluno04n1 + aluno04n2 + aluno04n3) / 3

	// c. Em sua função main imprima:
	// c.1- "Os alunos da turma são:"
	fmt.Println()
	fmt.Println("Os alunos da turma são:", aluno01, ",", aluno02, ",", aluno03, "e", aluno04)
	fmt.Println()

	// c.2- Na linha seguinte imprima o valor das variáveis com os nomes dos 4 alunos.
	fmt.Println("Aluno-1:", aluno01, "notas", aluno01n1, ",", aluno01n2, "e", aluno01n3)
	fmt.Println("Aluno-2:", aluno02, "notas", aluno02n1, ",", aluno02n2, "e", aluno02n3)
	fmt.Println("Aluno-3:", aluno03, "notas", aluno03n1, ",", aluno03n2, "e", aluno03n3)
	fmt.Println("Aluno-4:", aluno04, "notas", aluno04n1, ",", aluno04n2, "e", aluno04n3)
	fmt.Println()

	// c.3- Chame uma nova função chamada media
	media()
	diplomas()
	reprovados()

}

func media() {

	// d. a função media deverá imprimir:
	// d.1- A média final de cada aluno indicando o nome do mesmo;
	// d.2- Um resultado booleano que represente "true" para cada aluno que foi aprovado (cuja média seja igual ou maior a 7.0)
	fmt.Println("Aluno-1: ", aluno01, "media: ", media01)
	fmt.Println("A média do Aluno-1", aluno01, "foi", math.Round(media01*100)/100)
	fmt.Println("Aluno-1: ", aluno01, "ficou acima da média?", media01 >= 7)

	fmt.Println("Aluno-2: ", aluno02, "media: ", media02)
	fmt.Println("A média do Aluno-2", aluno02, "foi", math.Round(media02*100)/100)
	fmt.Println("Aluno-1: ", aluno02, "ficou acima da média?", media02 >= 7)

	fmt.Println("Aluno-3: ", aluno03, "media: ", media03)
	fmt.Println("A média do Aluno-3", aluno03, "foi", math.Round(media03*100)/100)
	fmt.Println("Aluno-1: ", aluno03, "ficou acima da média?", media03 >= 7)

	fmt.Println("Aluno-4: ", aluno04, "media: ", media04)
	fmt.Println("A média do Aluno-4", aluno04, "foi", math.Round(media04*100)/100)
	fmt.Println("Aluno-1: ", aluno04, "ficou acima da média?", media04 >= 7)

}

func diplomas() {
	// e. Crie uma nova função chamada diplomas
	// e.1- A função main deverá chamar também a função diplomas.
	// e.2- A função diplomas irá imprimir a seguinte linha para cada aluno que foi aprovado:
	// "Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana <NOME>."
	//diplomas =
	fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana,", aluno03)
	fmt.Println("Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana,", aluno04)

}

func reprovados() {
	// f. Crie uma função chamada reprovados.
	// f.1- A função main também deverá chamar a função reprovados
	// f.2- A função reprovados deverá imprimir o seguinte para cada aluno que foi reprovado:
	// "Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima <ALUNO>."
	//reprovados
	fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima,", aluno01)
	fmt.Println("Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima,", aluno02)

}
