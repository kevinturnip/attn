package model

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type User struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func AllUsers(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var users []User
		db.Find(&users)
		fmt.Println("{}", users)

		return c.JSON(http.StatusOK, users)
	}
}

func NewUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		db.Create(&User{Email: email, Password: password})
		return c.String(http.StatusOK, email+" user successfully created")
	}
}

func DeleteUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		email := c.Param("email")

		var user User
		db.Where("email = ?", email).Find(&user)
		db.Delete(&user)

		return c.String(http.StatusOK, email+" user successfully deleted")
	}
}

func GetUser(db *gorm.DB, email string, password string) User {
	var user User
	db.Where("email = ?", email, "password = ?", password).Find(&user)
	return user
}

func UpdateUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.Param("name")
		email := c.Param("email")
		var user User
		db.Where("name=?", name).Find(&user)
		user.Email = email
		db.Save(&user)
		return c.String(http.StatusOK, name+" user successfully updated")
	}
}
