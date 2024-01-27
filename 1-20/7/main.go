package main

import (
	"fmt"
)

// estudo do map
func main() {
	salarios := map[string]int{"Teste":1000, "Teste2":2000}
	delete(salarios, "Teste")
	salarios["Tes"] = 5000

	// formas de declarar
	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Wesly"] = 1000

	// blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}