package main

import "fmt"

func array1() {
  arreglo := [3]string{
    "Alex",
    "Pablo",
    "Juan",
  }

  // Este arreglo [Alex Pablo Juan] tiene 3 elementos
  fmt.Printf("Este arreglo %v tiene %v elementos\n", arreglo, len(arreglo))

  arreglo[1] = "Hernan"
  // Cambiar el valor de un elemento
  for i:= 0; i < len(arreglo); i++ {
    fmt.Print(arreglo[i], " ") // Alex Hernan Juan
  }
  fmt.Print("\n")
}

func array2() {
  var matriz [2][2] int
  fmt.Println(matriz)
}


func main() {
  array1()
  array2()
}
