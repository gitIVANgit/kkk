package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title aaaaaaaaaaaaaaaaa
// @version 1.0
// @description kkkkkkkkkkkke server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name kkkkkkkkkkk
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name kkkk
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
