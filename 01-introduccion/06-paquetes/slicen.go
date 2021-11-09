package main
import "fmt"

func main() {
  // slicen
  numeros := []int{1, 2, 3}
  fmt.Println(numeros, len(numeros))
  // agregar datos

  numeros = append(numeros, 4)
  numeros = append(numeros, 7)
  fmt.Println(numeros, len(numeros))

}
