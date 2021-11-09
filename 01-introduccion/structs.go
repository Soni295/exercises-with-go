package main

import "fmt"

type User struct{
  edad int
  nombre, apellido string
}

func main() {

  juan := User{}
  fmt.Println(juan) // {0   }

  pablo := User{nombre:"Pablo", apellido:"Ramidez"}
  fmt.Println(pablo) // {0 Pablo Ramidez}

  manuel := User{10, "Manuel", "Fernandez"}
  fmt.Println(manuel) // {10 Manuel Fernandez}

  uriel := new(User) // puntero
  fmt.Println(uriel) // &{0    }

  uriel.nombre = "Uriel"
  fmt.Println(uriel) // &{0 Uriel }
}
