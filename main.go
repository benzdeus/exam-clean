package main

import (
	"demo-clean/db"
	"demo-clean/route"

	"github.com/labstack/echo/v4"
)

func main() {

	dbPG, tearDown := db.New()
	defer tearDown()

	e := echo.New()

	route.InitRoute(e, dbPG)

	e.Start(":3000")

}
