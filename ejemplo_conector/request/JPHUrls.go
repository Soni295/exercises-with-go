package request

import (
  "net/http"
  "strconv"
)

const (
  urlBase = "https://jsonplaceholder.typicode.com/"
  urlTodos = urlBase + "todos/"
)

type ShortPost struct {
  // unint16 - 16-bits(0 a 65535)
  UserId uint16 `json:"userId"`
  Id uint16 `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

func Get_JPH_by_id(id int)(data *ShortPost, err error){
  url := urlTodos + strconv.Itoa(id)
  res, err := http.Get(url)
  if err != nil { return nil, err }

  defer res.Body.Close()

  data, err = Convert_respose_data(res)

  if err != nil { return nil, err }
  return data, nil
}
