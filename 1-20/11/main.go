package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// Estudo das Structs
// Não conseguimos utilizar herança mas podemos trabalhar com composição
func main() {
	osni := Cliente{
		Nome:  "Osni Antônio Küchler",
		Idade: 40,
		Ativo: true,
	}

	// posso alterar os valores
	osni.Idade = 41

	fmt.Println(osni.Idade)
}
