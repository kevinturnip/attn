package model

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	Name  string `json:"Name"`
	jwt.StandardClaims
}

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// func Login(db *gorm.DB) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		var u User
// 		// if err := c.Bind(&u); err != nil {

// 		// 	return c.String(http.StatusNotAcceptable, "Invalid json provided")
// 		// }
// 		u.Email = c.Param("email")
// 		u.Password = c.Param("password")

// 		//compare the user from the request, with the one we defined:
// 		user := GetUser(db, u.Email, u.Password)
// 		// if u.Name != u.Name || user.Password != u.Password {
// 		// 	c.JSON(http.StatusUnauthorized, "Please provide valid login details")
// 		// 	return
// 		// }
// 		if user.ID == 0 {
// 			return c.String(http.StatusNotAcceptable, "email/password not valid")
// 		}
// 		token, err := CreateToken(uint64(user.ID))
// 		if err != nil {

// 			return c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		}
// 		return c.JSON(http.StatusOK, token)
// 	}
// }

// func Login(c echo.Context, db *gorm.DB) error {
func Login(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		// // Throws unauthorized error
		// if username != "jon" || password != "shhh!" {
		// 	return echo.ErrUnauthorized
		// }
		log.Println("masuk sini ga sihhhhhhhhhh")
		log.Println(email, password)
		user := GetUser(db, email, password)
		// if u.Name != u.Name || user.Password != u.Password {
		// 	c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		// 	return
		// }
		if user.ID == 0 {
			return c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		}

		// Set custom claims
		claims := &JwtCustomClaims{
			user.Email,
			user.Name,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
