package main

import (
	"fmt"
)

var y = 75

// Sprint, Sprintln, Sprintf

func main() {
	fmt.Println(y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)

	fmt.Println(s)

}
