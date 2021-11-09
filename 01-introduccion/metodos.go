package main
import "fmt"

type User struct{
  edad int
  nombre, apellido string
}

func (this User) nombre_completo() string {
  return this.nombre + " " + this.apellido
}

func (this *User) set_name(nombre string) { // se pasa el puntero para modificar el valor en caso contrario no lo modifica por que es una copia
  this.nombre = nombre
}

func main() {
  uriel := new(User)
  uriel.nombre = "Uriel"
  uriel.apellido = "Hernandez"
  fmt.Println(uriel.nombre_completo()) // Uriel

  uriel.set_name("Martin")

  fmt.Println(uriel.nombre_completo()) // Uriel
}
