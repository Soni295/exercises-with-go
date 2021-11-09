package main
import "net/http"

func main() {

  fileServer := http.FileServer(http.Dir("public"))

  http.Handle("/", http.StripPrefix("/", fileServer))

 // http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}

/*
func handler(w http.ResponseWriter,r *http.Request){
  http.ServeFile(w, r, r.URL.Path[1:])
}
*/
