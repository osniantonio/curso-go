package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	nomeArquivo := "arquivo.txt"

	// Criando um arquivo
	f, err := os.Create(nomeArquivo)
	if err != nil {
		panic(err)
	}

	// Escrevendo dados no arquivo
	tamanhoString, err := f.WriteString("Hello, world\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanhoString)

	tamanhoBytes, err := f.Write([]byte("Escrevendo dados no arquivo\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanhoBytes)

	f.Close()

	// Leitura
	arquivo, err := os.ReadFile(nomeArquivo)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(arquivo))

	// Leitura de pouco em pouco para arquivos muito grandes - Gigas Bytes...
	arquivo2, err := os.Open(nomeArquivo)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Removendo um arquivo
	err = os.Remove(nomeArquivo)
	if err != nil {
		panic(err)
	}
}
