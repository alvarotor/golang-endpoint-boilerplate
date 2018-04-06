package users

import (
	"golang-endpoint-boilerplate/db" // --> Change this directory path to the name of your folder / project

	"golang.org/x/crypto/bcrypt"
)

func ReadAll() []db.User {
	return db.UsersAll()
}

func CreateUser(user *db.User) {
	user.Password = hashPassword(user.Password)
	db.UsersCreate(user)
}

func ReadUser(id string) (db.User, error) {
	user, err := db.UserRead(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func ReadUserByEmail(email string) (db.User, error) {
	user, err := db.UserReadByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func IsAdmin(email string) bool {
	user, err := db.UserReadByEmail(email)
	if err != nil {
		return false
	}
	if user.Admin == true {
		return true
	} else {
		return false
	}
}

func Authenticate(email string, password string) bool {
	user, err := db.UserReadByEmail(email)
	if err != nil {
		return false
	}
	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errHash == nil {
		return true
	} else {
		return false
	}
}

func hashPassword(password string) string {
	passByte := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} else {
		return string(hashedPassword)
	}
}
