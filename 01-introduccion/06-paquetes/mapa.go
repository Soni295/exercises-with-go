package main
import "fmt"


func main() {
  dias := make(map[int]string)
  fmt.Println(dias) // map[]

  dias[1] = "Domingo"
  dias[2] = "Lunes"
  dias[3] = "Martes"
  dias[4] = "Miercoles"
  dias[5] = "Jueves"
  dias[6] = "Viernes"
  dias[7] = "Sabado"

  fmt.Println(dias)
  // map[1:Domingo 2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:Sabado]

  delete(dias, 1)
  fmt.Println(dias)
  // map[2:Lunes 3:Martes 4:Miercoles 5:Jueves 6:Viernes 7:Sabado]

  estudiantes := make(map[string][]string)

  estudiantes["Primero"] = []string{"Mario", "Hernesto", "Juan"}
  estudiantes["Segundo"] = []string{"Marta", "Maria", "Pedro"}
  fmt.Println(estudiantes)
  //map[Primer:[Mario Hernesto Juan] Segundo:[Marta Maria Pedro]]
  fmt.Println(estudiantes["Primero"][0])
}
