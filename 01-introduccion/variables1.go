package main

import (
  "fmt"
  "strconv"
  "math/cmplx"
  "math"
)

func declaracionDeVariables(){
  var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
  )

  var nombre1 string
  nombre1 = "Martin"
  var i, j int = 1, 2

  fmt.Println(i, j) // 1 2
  var nombre2 string = "Juan"

  // := asume el tipo de variable segun el valor que le asigno
  edad := 26
  fmt.Println(nombre1, "y", nombre2, "tienen", edad, "Años")

  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}




func aritmetica() {
  a := 30
  b := 20
  result := a + b

  fmt.Println("Suma:", result)

  result = a - b
  fmt.Println("Resta:", result)

  result = a * b
  fmt.Println("Multiplicacion:", result)

  div := 30.0 / 20.0
  fmt.Println("Division:", div)

  result = a % b
  fmt.Println("Modulo:", div)
}


func pkg_print(){
  fmt.Print("Hola ")
  fmt.Print("me llamo Pepe")
  // Hola me llamo Pepe

  dia := 05
  mes := "Noviembre"
  fmt.Printf("\nSabias que hoy es %d de %s", dia, mes)

  // Hoy es 5 de Noviembre

  var nombre string
  fmt.Print("\n¿Comó te llamas?\n-> ")
  // sirve de input
  fmt.Scanln(&nombre)
  fmt.Printf("Ah Mucho gusto %s es un placer conocerte", nombre)
}

/*
  fmt.Print =>
    muestra en consola sin salto de linea
  fmt.Printf =>
    permite insertar los valores de las variables
    en la salida por consola en string
    example
      hola := "hola"
      mundo := "mundo\n"
      fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)
    %s = string
    %d = number
    %v = any
  fmt.Sprintf => retorna un string con injeccion de variables
    example
      mensaje := fmt.Sprintf("Hola, %v y su edad es %v \n", nombre, edad)
  fmt.Scanln(arg) toma como argumento una variable en la que
  metera el valor decidido
*/





func saludar(nombre string) {
  fmt.Println("Hola", nombre)
}
func sumar(num1, num2 int) int {
  return num1 + num2
}
func swap(x, y string) (string, string) {
  return y, x
}
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}
func funciones() {
  saludar("Hernan")
  fmt.Println(sumar(10, 30))
  a, b := swap("hello", "world")
  fmt.Println(a, b)
  fmt.Println(split(17))
}

func conversiones() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)

  fmt.Println(x, y, z, f)
  fmt.Printf("f is of type %T\n", f)
  edadNum := 24
  edadStr := strconv.Itoa(edadNum)
  fmt.Println("tengo "+ edadStr)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func numbersMinMax() {
  const Big = 1 << 100
  const Small = Big >> 99

  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}


func main() {
  // pkg_print()
  // aritmetica()
  // valoresPorDefault()
  // funciones()
  // conversiones()
  // numbersMinMax()
}
