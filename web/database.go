package main

import (
  "os"
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

type User struct {
  ID    string
  EMAIL string
}

func database() {
  db, err := sql.Open("postgres", "host="+os.Getenv("PGHOST")+" port=5432 user="+os.Getenv("PGUSER")+" dbname="+os.Getenv("PGDATABASE")+" password="+os.Getenv("PGPASSWORD"))
  defer db.Close()

  if err != nil {
    panic(err)
  }

  // INSERT
  var empID string
  id := 3
  number := 4444
  err = db.QueryRow("INSERT INTO employee(emp_id, emp_number) VALUES($1,$2) RETURNING emp_id", id, number).Scan(&empID)

  if err != nil {
      fmt.Println(err)
  }

  fmt.Println(empID)

  // SELECT
  rows, err := db.Query("SELECT * FROM employee")

  if err != nil {
      fmt.Println(err)
  }

  var es []User
  for rows.Next() {
      var e User
      rows.Scan(&e.ID, &e.EMAIL)
      es = append(es, e)
  }
  fmt.Printf("%v", es)
}
