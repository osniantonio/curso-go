package main

import "fmt"

// Estudo Type assertation
func main() {
	var minharVar interface{} = "Wesley Willians"
	println(minharVar.(string))
	res, ok := minharVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
	res2 := minharVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}