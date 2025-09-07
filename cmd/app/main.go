// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
    "fmt"
    "github.com/seu-usuario/meu-projeto-go/internal/hello"
    "github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
)

// Função principal do programa
func main() {
    fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
    hello.SayHello()
	fmt.Print("O 7º número de fibonacci é: ", fibonacci.Fibonacci(7))
}
