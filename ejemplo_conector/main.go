package main
import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  //"encoding/json"
  "github.com/Soni295/ExampleJPH/database"
  "github.com/Soni295/ExampleJPH/request"
  "sync"
)

var wg sync.WaitGroup

func connec_save(id int){
  defer wg.Done()

  data, err :=request.Get_JPH_by_id(id)
  if err != nil {
    fmt.Printf("Error con la peticion nº%v\n Error: -> %v\n", id, err)
    return
  }
  err = database.Insert_short_post(data)

  if err != nil {
    fmt.Printf("\nError insertando: %v\n", err)
    return
  }else{
    fmt.Printf("Insertado correctamente endpoint -> %v\n", id)
  }
}

func main() {
  db, err :=  database.ConnecBD()
  if err != nil { return }
  defer db.Close()
  fmt.Printf("Conectado correctamente\n")

  ids := 20
  wg.Add(ids)
  for i := 1; i <= ids; i++ {
    go connec_save(i)
  }

  wg.Wait()
  fmt.Printf("Programa finalizado \n")
}
