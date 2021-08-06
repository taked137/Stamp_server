package main

import (
    "time"
    "math/rand"
    "net/http"

    "github.com/labstack/echo/v4"
)

type (
    beaconRequest struct {
        Quiz int `json:"quiz"`
        Beacon []int `json:"beacon"`
    }
    beaconResponse struct {
        ID int `json:"id"`
        Quiz int `json:"quiz"`
        URL string `json:"url"`
    }
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
    e.POST("/stamp/image", func(c echo.Context) error {
        request := new(beaconRequest)
        if err := c.Bind(request); err != nil {
            return err
        }

        // num := 0
        // for _,v := range request.Beacon {
        //     num += v
        // }

        rand.Seed(time.Now().UnixNano())
        num := rand.Intn(3)

        response := beaconResponse {
            ID: num,
            Quiz: request.Quiz,
            URL: "https://1.bp.blogspot.com/-3XfMA0UhT70/XlyfrsiokjI/AAAAAAABXn8/j_CLCc73TTEi-PCK19hnUwY3D-pJgmjvQCNcBGAsYHQ/s1600/drink_beer_yukata_man.png",
        }
        return c.JSON(http.StatusOK, response)
    })
    e.Logger.Fatal(e.Start(":1323"))
}
