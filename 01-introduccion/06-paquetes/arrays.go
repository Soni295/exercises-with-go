package main
import "fmt"

func main() {
  var numeros [5]int
  fmt.Println(numeros)

  numeros[0] = 10
  numeros[1] = 20
  numeros[2] = 30

  fmt.Println(numeros)
  fmt.Println(numeros[4])

  nombres := [3]string{"Alex", "Roel", "Juan"}
  fmt.Println(nombres)

  colores := [...]string{
    "Rojo",
    "Verd",
    "Negro",
    "Azul",
  }

  fmt.Println(colores, "tiene", len(colores))

  monedas := [...]string{
    0: "Dolar",
    2: "Soles",
    3: "Pesos",
    5: "Euro",
  }
  fmt.Println(monedas, "tiene", len(monedas))

  subMoneda := monedas[:3]
  fmt.Println(subMoneda)
}
