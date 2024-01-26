package main

// Estudo Ponteiros e Structs
type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

// Ponteiros e Structs
func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}