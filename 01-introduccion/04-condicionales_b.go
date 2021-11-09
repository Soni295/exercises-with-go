package main

import "fmt"

/*
  && => AND
  || => OR
*/

func main() {
  var dinero uint8 = 50
  var credito uint8 = 100
  var cupo bool = true

  if dinero != 0 || credito != 0 {
    fmt.Println("Puedes comprar aqui")
    if cupo {
      fmt.Println("Como tiene el cupon te haremos un descuento")
    } else if dinero + credito > 20 {
      fmt.Println("Podrias conseguir un cupon si lo deseas")
    } else {
      fmt.Println("Es una lastima que no tengas cupon")
    }
  }
}
