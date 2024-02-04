package main

import (
	"context"
	"fmt"
)

// por convenção as variáveis de contexto sempre vem primeio no fonte
// exemplo de uso: passando informações utilizadas em métricas ou Correlation ID
// não é muito utilizada mas importante conhecer
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
