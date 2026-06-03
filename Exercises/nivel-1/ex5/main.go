package main

import (
	"fmt"
)

type dog int

var x dog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println("↑ foi x.\n↓ é y!") /*Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"*/
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
