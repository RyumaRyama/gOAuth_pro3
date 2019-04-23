package app

import (
   "fmt"
   "os"
   "github.com/jinzhu/gorm"
   _ "github.com/lib/pq"
)

type User struct {
   ID int64 `gorm:"primary_key" json:"id"`
   Name string `json:"name"`
}

type Users []User

func Add_user(email string) {
   db, err := gorm.Open("postgres", "user="+os.Getenv("POSTGRES_USER")+" password="+os.Getenv("POSTGRES_PASSWORD")+" dbname="+os.Getenv("POSTGRES_DB")+" sslmode=disable")
   if err != nil {
      panic(err)
   }
   defer db.Close()
   db.AutoMigrate(User{})
   var user = User{Name: email}
   db.NewRecord(user)
   db.Create(&user)
   db.Save(&user)

   var users = Users{}
   db.Find(&users) // SELECT * FROM users;
   fmt.Println(users)
}
