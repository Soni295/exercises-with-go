package main
import "fmt"

func main() {
  /*
    1. Es una direccion en memoria.
    2. En lugar de tener el valor,
       tenemos la direccion en la que esta el valor.
    3. X, Y => 0xc000018130 => 6
    4. X => 0xc000018130 => 5
    5. Y = 5
    *T => *int, *string, *Struct
    Valor del puntero por defecto es nil
  */

  var x, y *int
  entero := 10
  x = &entero
  y = &entero


  fmt.Printf("La posicion en memoria de x es %v\n", x)
  fmt.Printf("La posicion en memoria de y es %v\n", y)

  fmt.Printf("El valor de x es = %v\n", *x)
  fmt.Printf("El valor de y es = %v\n", *y)

  *x = 5

  fmt.Printf("El valor de x es = %v\n", *x)
  fmt.Printf("El valor de y es = %v\n", *y)
}
