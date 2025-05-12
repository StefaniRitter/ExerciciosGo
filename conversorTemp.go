package main
import "fmt"

func main() {
    var fahren, celsius float32
    fmt.Println("Informe uma temperatura em Fahrenheit: ")
    fmt.Scan(&fahren)
    
    celsius = (fahren - float32(32)) * 5
    celsius = celsius/9
    
    fmt.Printf("%.2f graus em Fahrenheit equivalem a %.2f graus Celsius", fahren, celsius)
 
}