package model

import (
	"attn/util"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Currencies struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ConversionRate struct {
	Id             int     `json:"id"`
	CurrencyIdFrom int     `json:"currency_id_from"`
	CurrencyIdTo   int     `json:"currency_id_to"`
	Rate           float64 `json:"rate"`
}

func AllCurrencies(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var currencies []Currencies
		db.Find(&currencies)
		fmt.Println("{}", currencies)

		return c.JSON(http.StatusOK, currencies)
	}
}

func NewCurrencies(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		log.Println("masuk ga sih")
		var resp util.Response
		id := c.FormValue("id")
		intId, _ := strconv.Atoi(id)
		name := c.FormValue("name")
		curr := Currencies{Id: intId, Name: name}
		currency, err := CreateNewCurrency(curr, db)
		if err != nil {
			resp = util.GenerateResp(currency, "Failed create currency")
		} else {
			resp = util.GenerateResp(currency, "Success Create Currency")
		}
		return c.JSON(http.StatusOK, resp)
	}
}

func AllConversionRate(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var conversionRate []ConversionRate
		db.Find(&conversionRate)
		fmt.Println("{}", conversionRate)

		return c.JSON(http.StatusOK, conversionRate)
	}
}

func NewConvertedCurrency(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		from := c.FormValue("from")
		idFrom, err := strconv.Atoi(from)
		to := c.FormValue("to")
		idTo, err := strconv.Atoi(to)
		rate := c.FormValue("rate")
		rateFloat, err := strconv.ParseFloat(rate, 64)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "bad request field rate")
		}

		conversionRate := &ConversionRate{CurrencyIdFrom: idFrom, CurrencyIdTo: idTo, Rate: rateFloat}
		conRate := GetConvertedCurrency(idFrom, idTo, db)
		if conRate.Id != 0 {
			resp := util.GenerateResp(conRate, "data already exists")
			return c.JSON(http.StatusBadRequest, resp)
		}
		db.Create(conversionRate)
		// var conversionRate ConversionRate
		// db.Where("currency_id_from = ? and currency_id_to", from, to).Find(&conversionRate)
		// if conversionRate.Rate != 0 {
		// 	result = amountFloat / conversionRate.Rate
		// }
		// res := fmt.Sprintf("Result: %f", result)
		return c.JSON(http.StatusOK, conversionRate)
	}
}

func GetConvertedCurrencyRate(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var result float64
		var res string
		var resp util.Response
		from := c.FormValue("from")
		to := c.FormValue("to")

		amount := c.FormValue("amount")
		amountFloat, err := strconv.ParseFloat(amount, 64)
		log.Println(err)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "bad request field amount")
		}
		var conversionRate, conversionRate2 ConversionRate
		db.Where("currency_id_from = ? and currency_id_to", from, to).Find(&conversionRate)
		db.Where("currency_id_from = ? and currency_id_to", to, from).Find(&conversionRate2)
		if conversionRate.Id != 0 {
			result = amountFloat / conversionRate.Rate
		} else if conversionRate2.Id != 0 {

			result = amountFloat / conversionRate2.Rate

		}
		if result != 0 {
			res = fmt.Sprintf("Result: %.2f", result)
			resp = util.GenerateResp(res, "Success get Rate")
		} else {
			resp = util.GenerateResp(0, "Bad Request")
		}
		return c.JSON(http.StatusOK, resp)
	}
}

func GetConvertedCurrency(from, to int, db *gorm.DB) (conRate ConversionRate) {
	db.Where("currency_id_from = ? and currency_id_to", from, to).Find(&conRate)
	return
}

func CreateNewCurrency(curr Currencies, db *gorm.DB) (Currencies, error) {

	db = db.Create(&curr)

	if db.Error != nil {
		return Currencies{}, db.Error
	}
	return curr, nil
}

func MockCreateCurrency(name string, id int) (curr Currencies, err error) {
	curr.Id = id
	curr.Name = name
	if reflect.TypeOf(id).Kind() != reflect.Int {
		err = errors.New("not Integer data type")
		return
	}
	if reflect.TypeOf(name).Kind() != reflect.String {
		err = errors.New("not String data type")
		return
	}
	return
}
