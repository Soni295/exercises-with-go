package main
import "fmt"

/*
  Declarar variables
  var variable tipo = calor
  variable := valor // asume y le asigna el tipo de valor a la variable
  var variable1, variable2 tipo = valor1, valor2
  var (
    variable tipo = valor
    variable tipo = valor
    variable tipo = valor
  )

  las variables tienen un valor por defecto
  string " "
  number 0
  array []
  bool false
*/

func main() {
  var num1, num2 int = 3, 4
  var num1_float float64 = float64(num1)
  fmt.Printf("num1 = %d type %T\n", num1, num1)
  fmt.Printf("num2 = %d type %T\n", num2, num2)
  fmt.Printf("num1_float = %v type %T\n", num1_float, num1_float)
}
