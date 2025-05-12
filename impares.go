package main
import "fmt"

func main() {
    var n, i, cont uint
    fmt.Println("Quantos ímpares você deseja ver? ")
    fmt.Scan(&n)
    
    fmt.Printf("Os primeiros %d ímpares são:\n", n)
    for i = 1; cont < n; i+=2 {
        fmt.Println(i)
        cont += 1
    }
 
}