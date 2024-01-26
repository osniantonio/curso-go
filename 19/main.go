package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}

	numerosII := []string{"um", "dois", "três"}
	for _, v := range numerosII {
		println(v)
	}

	for {
		println("Hello, world!")
	}
}
