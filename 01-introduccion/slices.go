package main
import "fmt"

func slice1(){
  slice := []int{1, 2, 3, 4} // slice
  fmt.Println(slice) // [1 2 3 4]
}

func slice2(){
  arreglo := [3]int{1, 2, 3} // array
  slice := arreglo[:2] //slicing
  fmt.Println(slice) // [1 2]
}

func slice3(){
  // make crea un data, en este caso el primer parametro
  // es el tipo
  // el segundo la cantidad
  // y el tercero la capasidad

  slice := make([]int, 3, 5)
  fmt.Println(slice) // [0 0 0]
  fmt.Printf("El slice %v tiene %v de elementos y una capasidad de %v elementos\n", slice, len(slice), cap(slice))
  // El slice [0 0 0] tiene 3 de elementos y una capasidad de 5 elementos

  slice = append(slice, 2)
  slice = append(slice, 3)
  fmt.Println(slice) // [0 0 0 2 3]
  fmt.Printf("El slice %v tiene %v de elementos y una capasidad de %v elementos\n", slice, len(slice), cap(slice))
  // El slice [0 0 0 2 3] tiene 5 de elementos y una capasidad de 5 elementos

  slice = append(slice, 4)
  fmt.Println(slice) // [0 0 0 2 3 4]
  fmt.Printf("El slice %v tiene %v de elementos y una capasidad de %v elementos\n", slice, len(slice), cap(slice))
  // El slice [0 0 0 2 3 4] tiene 6 de elementos y una capasidad de 10 elementos

  // cada vez que el slice desborda su capasidad se crea un nuevo array con el doble
  // de capasidad
}


func main() {
  /*
  Los slices tienen:
    Puntero al arreglo
    Longitudad del arreglo al que apunta
    Capacidad
  */

  slice1()
  slice2()
  slice3()
}
