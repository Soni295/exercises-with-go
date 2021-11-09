package main

import "fmt"

/*
  if(condicion) { // el parentecis es opcional
    si se cumple haz esto
  } else {
    y sino haz esto otro
  }
*/

func main() {
  var edad int

  edad = 15
  fmt.Printf("Tienes %d años ", edad)

  if edad >= 18 {
    fmt.Print("asi que puedes pasar\n")
  } else {
    fmt.Print("necesitas tener al menos 18 años para entrar\n")
  }
  // Tienes 15 años necesitas tener al menos 18 años para entrar

  edad = 28
  fmt.Printf("Tienes %d años ", edad)

  if edad >= 18 {
    fmt.Print("asi que puedes pasar\n")
  } else {
    fmt.Print("necesitas tener al menos 18 años para entrar\n")
  }
  // Tienes 28 años asi que puedes pasar
}
