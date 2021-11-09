package main
import "fmt"

/*
  + => suma
  - => resta
  * => multiplicacion
  / => divicion
  % => modulo de divicion(resto)
*/

func main() {
  a := 30
  b := 20

  result := a + b
  fmt.Printf("La suma de %d y %d es igual a %d\n", a, b, result)

  result = a - b
  fmt.Printf("La resta de %d y %d es igual a %d\n", a, b, result)

  result = a * b
  fmt.Printf("La multiplicación de %d y %d es igual a %d\n", a, b, result)

  result = 30 / 20
  modulo := a % b
  fmt.Printf("La divición de %d y %d es igual a %d con un modulo de %d\n", a, b, result, modulo)
}
