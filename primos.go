package main
import "fmt"

func main() {
    var (
        n, i uint
        primo bool = true
        )
    fmt.Println("Informe um número inteiro e positivo para verificar se é primo: ")
    fmt.Scan(&n)
    
    for i = 2; i < n; i++ {
        if n%i == 0 {
            primo = false
        }
    }
    
    if primo && n > 1 {
        fmt.Printf("O número %d é primo.", n)
    } else {
        fmt.Printf("O número %d não é primo.", n)
    }
    
}