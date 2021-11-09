package main
import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  ejecucion := read_file()
  fmt.Println(ejecucion)
}

func read_file() bool{
  archivo, err := os.Open("./hola.txt")
  defer func() {
    archivo.Close()
    fmt.Println("se cerro")
  }()

  if err != nil {
    fmt.Println("Hubo un error")
  }
  scanner := bufio.NewScanner(archivo)

  i := 1
  for scanner.Scan() {
    linea := scanner.Text()
    fmt.Println(i,"|",  linea)
    i++
  }
  return true
}
