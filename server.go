package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.GET("/:user", func(c echo.Context) error {
        user := c.Param("user")
        return c.String(http.StatusOK, "Hello, World!" + user)
    })
    e.Logger.Fatal(e.Start(":1323"))
}