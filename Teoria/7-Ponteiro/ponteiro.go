package main

import "fmt"

func main() {
	n1 := 10
	n2 := &n1

	fmt.Println(n1)
	fmt.Println(*n2)

	*n2 = 15

	fmt.Println(n1)
	fmt.Println(*n2)
}
