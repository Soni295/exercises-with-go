package main
import "fmt"

/*
  for initialization; condition; post{
    // todo
  }
*/

func main() {
  sum := 0

  for i:= 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum) // 45
}
