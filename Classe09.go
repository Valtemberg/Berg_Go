package main

import (
	"fmt"
)

var y = 32
var d = 12345687

func main() {
	fmt.Printf("Em hexadecimal com 0x %#x, em binario %b, e em hexadecimal %x\n", y, y, y)
	fmt.Printf("O valor de y é %v, já os demais valores são: %#x\t%b\t%x\n", y, y, y, y)
	fmt.Printf("O valor de d é %v, já os demais valores são: %#x\t%b\t%x\n", d, d, d, d)

}
