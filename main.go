package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	s := NewSimpleStore()
	e := echo.New()

	e.PUT("/v1/:key", getPutHandler(s))
	e.GET("/v1/:key", getGetHandler(s))
	e.DELETE("/v1/:key", getDeleteHandler(s))

	log.Fatal(e.Start("localhost:8090"))
}
