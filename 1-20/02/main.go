package main

// const - não consigo alterar o valor
const a = "Hello, world!"

// var - permite alterar o valor da variável
// variáveis declaradas como globais
// ao declarar o go já infere um valor para a variável
// se uma variável foi criada como bool somente valores booleanos são permitidos por ser fortemente tipada
var (
	b bool = true
	c int
	d string
	e float64
)

func main() {
	// o go não permite ter variáveis de escopo declaradas e não utilizadas
	// var a string = "x"

	// para declarar com shorthands usa-se := (somente utilizar os dois pontos na primeira vez que e quando for declarada)
	a := "teste"
	println(a)
}