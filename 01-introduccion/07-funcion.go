package main
import "fmt"

/*
  func nombre(parametros)(tipo de retorno){
      // function body.....
  }
*/

func main() {
  alto := 10
  ancho := 30
  fmt.Printf("El area de un rectangulo de %dm de alto y %d m de ancho es %dm\n", alto, ancho, area(alto, ancho))

  base := 30
  altura := 50
  fmt.Printf("El area de un triangulo de %dm de base y %d m de altura es %dm\n", base, altura, areaTriangulo(base, altura))
}
func area(alto, ancho int) int {
  return alto * ancho
}
func areaTriangulo(base, altura int)( areaT int){
  areaT = base * altura / 2
  return
}
