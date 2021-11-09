package main
import "fmt"

type User interface {
  Permissions() int // 1 - 5
  Name() string
}

type Admin struct {
  name string
}
func (this Admin) Permissions() int {
  return 5
}
func (this Admin) Name() string {
  return this.name
}

type Editor struct{
  name string
}

func (this Editor) Permissions() int {
  return 3
}
func (this Editor) Name() string {
  return this.name
}


func auth(user User) string {
  if user.Permissions() > 4 {
    return user.Name() + " tiene permisos"
  }
  return user.Name() + " no tiene permisos"
}

func main() {
  admin := Admin{"Juan"}
  editor := Editor{"Pablo"}
  fmt.Println(auth(admin)) // Juan tiene permisos
  fmt.Println(auth(editor)) // Pablo no tiene permisos

  usuarios := []User{Admin{"Uriel"}, Editor{"Fulano"}}

  for _, usuario := range usuarios {
    fmt.Println(usuario) // {Uriel} {Fulano}
  }
}
