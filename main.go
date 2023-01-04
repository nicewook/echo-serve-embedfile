package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed static
var static embed.FS

func main() {

	e := echo.New()

	e.GET("/csv", func(c echo.Context) error {
		p, err := static.ReadFile("static/userinfo.csv")
		if err != nil {
			log.Println(err)
			return err
		}
		c.Response().Header().Set(echo.HeaderContentDisposition, "attatchment; filename=userinfo.csv")
		return c.Blob(http.StatusOK, "text/csv", p)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
