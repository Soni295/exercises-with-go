package request
import (
  "encoding/json"
  "net/http"
)

func Convert_respose_data(res *http.Response)(data *ShortPost, err error){
  err = json.NewDecoder(res.Body).Decode(&data)
  if err != nil { return nil, err }
  return data, nil
}
