package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

// Quando usar ponteiros
func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	total := soma(&minhaVar1, &minhaVar2)
	println(total)
}