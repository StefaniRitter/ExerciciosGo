package main
import "fmt"

func main() {
    var (
        soma int = 1
        x int
        n uint
        )
    fmt.Println("Informe um número inteiro: ")
    fmt.Scan(&x)
    fmt.Println("Informe um número inteiro não-negativo: ")
    fmt.Scan(&n)
    
    var i uint
    for i = 1; i <= n; i++ {
        soma *= x
    }
    
    fmt.Printf("O resultado de %v elevado à %v é igual a %v.", x, n, soma)
}