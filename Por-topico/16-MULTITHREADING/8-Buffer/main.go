package main

func main() {
	// cuidar para não utilizar o buffer de forma indevida
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
