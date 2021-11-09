package database
import (
  "database/sql"
  "fmt"
)

type UserDB struct {
  user, password, host, db string
}

func (this *UserDB) get_sentence() string {
  sentence := fmt.Sprintf("%s:%s@%s/%s", this.user, this.password, this.host, this.db)
  return sentence
}

func ConnecBD() (db *sql.DB, e error){
  userDB := UserDB{
    user: "sion", // "root" default
    password: "123456", // "" default
    host: "tcp(127.0.0.1:3306)",
    db: "go_example",
  }

  db, err := sql.Open("mysql", userDB.get_sentence())

  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    fmt.Printf("Error conectando: %v", err)
    return nil, err
  }

  return db, nil
}
