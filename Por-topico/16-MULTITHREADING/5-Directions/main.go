package main

import "fmt"

// receive-only: diz que esse canal somente recebe dados (seta do lado direito)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// send-only: essa somente esvazia o canal (seta do lado esquerdo)
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
