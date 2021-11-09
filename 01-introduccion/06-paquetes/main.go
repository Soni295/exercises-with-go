package main
import (
  "fmt"
  "math"
  "math/rand"
)

func add(num1 int, num2 int) int {
  return num1 + num2
}

func add2(num1, num2 int) int {
  return num1 + num2
}

func swap(x, y string) (string, string) {
  return y, x
}


/*
  Esto permite declarar las variables
  de la funcion arriba
  pero los retorna automaticamente
*/

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Println(math.Pi)
  fmt.Println(add(42, 13))
  fmt.Println(add2(42, 13))

  a, b := swap("Hello", "World")

  fmt.Println(a, b)

  fmt.Println(split(17))
}
