package main
import "fmt"

/*
  for initialization; condition; post{
    // todo
  }
*/
// finite
func main() {
  sum := 0
  for {
    sum++
    if sum % 2 == 0 { continue }
    if sum > 10 { break }
    fmt.Print(sum, " ") // 1 3 5 7
  }
}
