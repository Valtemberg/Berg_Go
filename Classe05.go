package main

import "fmt"

var y = 42
var z = "Sou uma variável string"

func main() {
	fmt.Println(y)        //ln = Line
	fmt.Printf("%T\n", y) // f = Format (Formatação) \n = Quebra de linha
	fmt.Println("E agora?")
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

//%T = Type / Tipo
