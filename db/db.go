package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Initialize
)

type User struct {
	gorm.Model
	Name       string
	UserName   string
	Email      string
	Password   string
	Admin      bool
	SuperAdmin bool
}

var db *gorm.DB
var err error

func initDB() {
	fmt.Println("Connecting to DB")
	var envVar = os.Getenv("db-connection-string")
	if envVar == "" {
		envVar = "host=localhost port=5432 user=gorm dbname=gorm password=local sslmode=disable"
	}
	db, err = gorm.Open("postgres", envVar)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
}

func InitialMigration() {
	initDB()
	defer db.Close()

	fmt.Println("Start initialMigration")

	db.DropTableIfExists(&User{})

	db.AutoMigrate(&User{})
	db.Create(&User{Name: "name1", UserName: "Username 1", Email: "email1@mail.com", Password: "$2a$10$zK5ZzxhAkXcgNYzU9ok8jeVor9oiPgR8U8dZ5bDoEgcCgglJ/KoRy"})
	db.Create(&User{Name: "name2", UserName: "Username 2", Email: "email2@mail.com", Password: "$2a$10$BG4qSC5.dzM96jqTVrzsMO98Y1P/fHUVmmER1l15d9iccjxxGoeBm", Admin: true})

	fmt.Println("Finished initialMigration")
}

func UsersAll() []User {
	initDB()
	defer db.Close()

	var users []User
	db.Find(&users)

	return users
}

func UsersCreate(user *User) {
	initDB()
	defer db.Close()

	db.Save(&user)
}

func UserRead(id string) (User, error) {
	initDB()
	defer db.Close()

	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UserReadByEmail(email string) (User, error) {
	initDB()
	defer db.Close()

	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
