package main
import (
  "net/http"
  "encoding/json"
)

type Curso struct {
  Title string `json:"title"`
  NumeroVideos int `json:"numeroVideos"`
}

type Cursos []Curso

func main() {
  server()
}

func server() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

    cursos := Cursos{
      Curso{"React.js", 3},
      Curso{"Vue.js", 3},
      Curso{"Angular.js", 3},
      Curso{"Node.js", 3},
    }
    json.NewEncoder(w).Encode(cursos)
  })
  http.ListenAndServe(":8000", nil)
}
