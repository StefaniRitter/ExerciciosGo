package main
import (
    "fmt"
    "bufio"
    "os"
    s "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    var (
        frase, palavra string
        tamFrase, tamPalavra, cont, i int
        igual bool
    )
    fmt.Println("Digite uma frase: ")
    frase, _ = reader.ReadString('\n')
    frase = s.TrimSpace(frase) //remove o \n do final

    fmt.Println("Qual palavra deseja procurar? ")
    fmt.Scan(&palavra)
    
    tamFrase = len(frase)
    tamPalavra = len(palavra)
    aparece := 0
    
    for cont = 0; cont <= (tamFrase - tamPalavra); cont ++ {
        igual = true
        for i = 0; i < tamPalavra; i++{
            if frase[cont+i] != palavra[i]{
                igual = false            
            }
        }
            
        if igual {
            aparece += 1
        }
    }
    
    fmt.Printf("A palavra %s aparece %d vezes na frase.", palavra, aparece)

}