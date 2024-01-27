package main

import "fmt"

// vazia implementa todo mundo
// type x interface {}

// Estudo das Interfaces Vazias
func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}