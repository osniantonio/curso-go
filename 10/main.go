package main

import (
	"fmt"
)

// estudo dos closures
func main() {
	total := func() int {
		return sum(1, 2, 3) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
