package main

import (
	"fmt"

	"github.com/osniantonio/goexpert/13-Packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
