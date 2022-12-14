package main

import (
	"fmt"
	"time"

	"github.com/aydinsercan/basketapi/controller"
	h "github.com/aydinsercan/basketapi/helper"
	db "github.com/aydinsercan/basketapi/model"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	//Read Config File
	if err := h.ReadConfig(); err != nil {
		panic(err)
	}

	// Create database if not exists
	if err := db.Migrate(); err != nil {
		fmt.Println("Migration Error :", err)
	}

	dbConfig := mysql.Config{
		User:                 h.Config.Database.User,
		Passwd:               h.Config.Database.Password,
		DBName:               h.Config.Database.DbName,
		Loc:                  time.Local,
		Net:                  fmt.Sprintf("tcp(%s:%d)", h.Config.Database.Host, h.Config.Database.Port),
		AllowNativePasswords: true,
	}

	// Open database connection
	if err := db.InitDB(dbConfig.FormatDSN()); err != nil {
		panic(err)
	}
	if err := db.Pool.Ping(); err != nil {
		panic(err)
	}
}
func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Set Routers
	controller.RegisterRouters(e.Group("/api"))

	// Start
	e.Start(fmt.Sprintf(":%d", h.Config.Port))
}
