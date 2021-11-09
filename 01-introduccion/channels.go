package main
import "fmt"

func main() {

  channel := make(chan string)


  go func(channel chan string){
    for {
      name := ""
      fmt.Scanln(&name)
      channel <- name
    }
  }(channel)

  msg := <- channel
  fmt.Println(msg)

  msg = <- channel
  fmt.Println(msg)
}
