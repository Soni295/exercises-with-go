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
  archivo, err := os.Open("./holaa.txt")
  defer func() {
    archivo.Close()
    fmt.Println("se cerro")
    recover()
  }()

  if err != nil {
    fmt.Println("Hubo un error")
    panic(err)
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
