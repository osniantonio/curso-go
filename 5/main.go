package main

import (
	"fmt"
)

// const - não consigo alterar o valor
const a = "Hello, world!"

// criando o seu próprio tipo
type ID int

// var - permite alterar o valor da variável
// variáveis declaradas como globais
// ao declarar o go já infere um valor para a variável
// se uma variável foi criada como bool somente valores booleanos são permitidos por ser fortemente tipada
var (
	b bool = true
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	// o go não permite ter variáveis de escopo declaradas e não utilizadas
	// var a string = "x"

	// para declarar com shorthands usa-se := (somente utilizar os dois pontos na primeira vez que e quando for declarada)
	// a := "teste"

	// criando um array com três posições fixas
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	
	for i, v := range meuArray {
		fmt.Printf("O valor do índice é %d e o valor é %d\n", i, v)
	}
}