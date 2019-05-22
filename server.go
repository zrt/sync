package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

var (
	currentname = "init"
	currenturl  = "https://yeah.moe/"
)

const urlprefix = "https://yeah.moe/"

type Response struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func pull(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{currentname, currenturl})
}
func push(c echo.Context) error {
	url := c.FormValue("url")
	name := c.FormValue("name")
	if url == "" {
		return c.String(http.StatusOK, "miss url")
	}
	if !strings.HasPrefix(url, urlprefix) {
		return c.String(http.StatusOK, "invalid url")
	}
	if name == "" {
		name = "Guest"
	}

	currenturl = url
	currentname = name

	return c.String(http.StatusOK, "ok")
}

func main() {
	e := echo.New()
	e.GET("/pull", pull)
	e.POST("/push", push)
	e.Logger.Fatal(e.Start("127.0.0.1:7003"))
}
