package database
import (
  "fmt"
  "github.com/Soni295/ExampleJPH/request"
)

type table_short_post struct {
  table_name, user_id, original_id, id, title, completed string
}

func (this table_short_post) create() string {
  return fmt.Sprintf("INSERT INTO %v(%v, %v, %v, %v) VALUES (?, ?, ?, ?)",
    this.table_name,
    this.user_id,
    this.original_id,
    this.title,
    this.completed,
  )
}

var table = table_short_post{
  table_name: "posts",
  user_id: "user_id",
  original_id:"original_id",
  id: "id",
  title:"title",
  completed: "completed",
}

func Insert_short_post(post *request.ShortPost) (err error){
  db, err :=  ConnecBD()
  if err != nil { return err }
  defer db.Close()
  sql_query, err := db.Prepare(table.create())
  defer sql_query.Close()

  if err != nil { return err }

  _, err = sql_query.Exec(post.UserId, post.Id, post.Title, post.Completed)

  if err != nil { return err }
  return nil
}
