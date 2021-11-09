package main
import "fmt"


type Human struct{
  name string
}

func (this Human) hablar() string {
  return  "Bla bla bla"
}

type Tutor struct{
  Human
}

func main() {
  tutor := Tutor{Human{"Uriel"}}
  fmt.Println(tutor.name) // Uriel
  fmt.Println(tutor.Human.name) // Uriel
  fmt.Println(tutor.hablar()) // Bla bla bla
}
