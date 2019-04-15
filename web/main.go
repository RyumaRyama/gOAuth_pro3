package main

import (
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  e.GET("/login", func(c echo.Context) error {
    return c.String(http.StatusOK, "hello")
  })
  e.Logger.Fatal(e.Start(":1124"))
}

