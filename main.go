package main

import (
	"attn/model"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func handlerFunc(msg string) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, msg)
	}
}

func handleRequest(db *gorm.DB) {
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// r := e.Group("/v1")

	// Configure middleware with the custom claims type
	// e.POST("/register", model.NewUser(db))
	// log.Println("sini masuk kan")
	// e.POST("/login", model.Login(db))
	// e.GET("/users", model.AllUsers(db))
	// config := middleware.JWTConfig{
	// 	Claims:     &model.JwtCustomClaims{},
	// 	SigningKey: []byte("secret"),
	// }
	// r.Use(middleware.JWTWithConfig(config))
	// // r.GET("", restricted)
	// // e.POST("/register/:email/:password", model.Login(db))

	// r.DELETE("/user/:name", model.DeleteUser(db))
	// r.PUT("/user/:name/:email", model.UpdateUser(db))

	//currency
	e.GET("/currencies", model.AllCurrencies(db))
	e.POST("/currency", model.NewCurrencies(db))

	// conversion rate
	e.POST("conversion_currency", model.NewConvertedCurrency(db))
	e.GET("conversion_currency", model.AllConversionRate(db))
	e.GET("converted_rate", model.GetConvertedCurrencyRate(db))

	e.Logger.Fatal(e.Start(":3000"))
}

func initialMigration(db *gorm.DB) {

	db.AutoMigrate(&model.User{}, model.Currencies{}, model.ConversionRate{})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db, err := gorm.Open("sqlite3", "sqlite3gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	initialMigration(db)
	handleRequest(db)
}
