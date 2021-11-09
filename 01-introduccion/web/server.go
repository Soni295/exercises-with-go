package main
import (
  "net/http"
  "io"
  "fmt"
)

func main() {
  http.HandleFunc("/holaMundo", func(w http.ResponseWriter, r *http.Request){
    fmt.Println("Hay una nueva peticion")
    io.WriteString(w, "Hola mundo")
  })
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}
func handler(w http.ResponseWriter,r *http.Request){
  fmt.Println("Hay una nueva peticion")
  io.WriteString(w, "Bienvenido al main")
}
