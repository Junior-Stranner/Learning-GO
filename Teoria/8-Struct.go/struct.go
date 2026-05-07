package main

import "fmt"

func main() {

	type cordenada struct {
		x int
		y int
	}

	cord := cordenada{1, 2}
	fmt.Println(cord)

}
