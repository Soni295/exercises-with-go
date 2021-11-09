package main
import (
  "fmt"
  "encoding/json"
)

type People struct {
  Number int `json: "number"`
}

func main() {
  text :=  `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`

  textBytes := []byte(text)
  people := People{}

  err := json.Unmarshal(textBytes, &people)

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(people)
}
