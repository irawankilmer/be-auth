package main

import (
	"be-blog/internal/bootstrap"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	bootstrap.InitAPP()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
