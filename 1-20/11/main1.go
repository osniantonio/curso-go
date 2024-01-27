package main

import (
	"fmt"
)

// Estudo das Interfaces
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// qualquer struct com esse método implementa essa interface (automaticamente)
// interface no go não permite passar propriedades e somente métodos
type Pessoa interface {
	Desativar()
}

type Cliente1 struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Teste 1",
		Idade: 30,
		Ativo: true,
	}
	Desativacao(wesley)
	wesley.Desativar()
	wesley.Endereco.Cidade = "São Paulo"
	fmt.Println(wesley)

	minhaEmpresa := Empresa()
	Desativacao(minhaEmpresa)
}
