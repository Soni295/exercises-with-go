package main

import "fmt"

func copiar(){
  slice := []int{1, 2, 3, 4}
  copia := make([]int,4)
  copia2 := make([]int,1)

  fmt.Println(slice) // [1 2 3 4]
  fmt.Println(copia) // [0 0 0 0]
  fmt.Println(copia2) // [0]

  //copy(destino, fuente)
  copy(copia, slice)
  copy(copia2, slice)

  fmt.Println(slice) // [1 2 3 4]
  fmt.Println(copia) // [1 2 3 4]
  fmt.Println(copia2) // [1]
}


func shortcut(){
  slice := []int{2, 4, 6, 8}
  copia := make([]int,len(slice), cap(slice)*2)

  fmt.Println(slice) // [2 4 6 8]
  fmt.Println(copia) // [0 0 0 0]

  copy(copia, slice)

  fmt.Println(slice) // [2 4 6 8]
  fmt.Println(copia) // [2 4 6 8]
}


func main() {
  /*
    La funcion copy copia
    el minimo de elementos
  */
  copiar()
  shortcut()
}
