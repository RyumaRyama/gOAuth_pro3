package main

import (
   "fmt"

   "github.com/jinzhu/gorm"
   _ "github.com/lib/pq"
)

type User struct {
   ID int64 `gorm:"primary_key" json:"id"`
   Name string `json:"name"`
}

type Users []User

func main() {
   db, err := gorm.Open("postgres", "user=e165742 password=pw dbname=gOAuth_pro sslmode=disable")
   if err != nil {
      panic(err)
   }
   defer db.Close()
   db.AutoMigrate(User{})
   var user = User{Name: "testname"}
   db.NewRecord(user)
   db.Create(&user)
   db.Save(&user)

   var users = Users{}
   db.Find(&users) // SELECT * FROM users;
   fmt.Println(users)
}
