package main

import (
	"fmt"
)

var y = 42
var z = `Sou uma variável "string"` //Crase podemos usanar quando queremos``
var a = `Berg disse
		"sou uma
Variável String"`
var k = "Daniel disse que era homem."

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	z = "43"
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(k)
	fmt.Printf("%T\n", k)
}

//Raw STRING
