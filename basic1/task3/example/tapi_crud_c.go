package main

import (
    "context"
    "fmt"
    "time"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID       uint `gorm:"primaryKey"`
    Name     string
    Age      int
    Birthday time.Time
}

func main() {
    db, err := gorm.Open(sqlite.Open("tapi_crud_c.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    if err := db.AutoMigrate(&User{}); err != nil {
        panic(err)
    }

    ctx := context.Background()

    user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
    result := db.WithContext(ctx).Create(&user)

    fmt.Println(user.ID)
    fmt.Println(result.Error)
    fmt.Println(result.RowsAffected)
}
