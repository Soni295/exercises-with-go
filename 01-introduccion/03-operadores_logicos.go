package main
import "fmt"

/*
  == -> igual a
  != -> diferente a
  <  -> menor que
  >  -> meyor que
  <= -> menor igual que
  >= -> meyor igual que
*/

func main() {
  a := 30
  b := 20
  var result bool

  result = a > b
  fmt.Printf("a(%d) es mayor que b(%d) -> %v\n", a, b, result)
  // a(30) es mayor que b(20) -> true

  result = a != b
  fmt.Printf("a(%d) es diferente que b(%d) -> %v\n", a, b, result)
  // a(30) es diferente que b(20) -> true

  result = a < b
  fmt.Printf("a(%d) es manor que b(%d) -> %v\n", a, b, result)

  // a(30) es manor que b(20) -> false
  result = a < b
  fmt.Printf("a(%d) es manor que b(%d) -> %v\n", a, b, result)
  // a(30) es manor que b(20) -> false

  result = a == b
  fmt.Printf("a(%d) es igual que b(%d) -> %v\n", a, b, result)
  // a(30) es igual que b(20) -> false
}
