package main
import "fmt"

/*
  switch declaracionOpcional; expressionOpcional{
    case exp: Statement..
    case exp: Statement..
    ...
    default: Statement..
  }
*/
func main() {
  switch dia := 3; dia { // Miercoles
    case 1:
    fmt.Println("Lunes")
    case 2:
    fmt.Println("Martes")
    case 3:
    fmt.Println("Miercoles")
    case 4:
    fmt.Println("Jueves")
    case 5:
    fmt.Println("Viernes")
    case 6:
    fmt.Println("Sabado")
    case 7:
    fmt.Println("Domingo")
    default:
    fmt.Println("Invalido")
  }
}
