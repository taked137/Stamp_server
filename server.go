package main

import (
    "time"
    "os"

    "math/rand"
    "net/http"
    "io/ioutil"

    "github.com/labstack/echo/v4"
    "github.com/google/uuid"
)

type (
    regulatoinResponse struct {
        Message string `json:"message"`
    }

    userRequest struct {
        Name string `json:"name"`
        Device string `json:"device"`
        Version string `json:"version"`
    }
    userResponse struct {
        UUID string `json:"uuid"`
    }
    
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
    e.POST("/user/create", func (c echo.Context) error {
        request := new(userRequest)
        if err := c.Bind(request); err != nil {
            return err
        }

        response := userResponse {}

        if(request.Name == "taked") {
            return c.JSON(http.StatusConflict, response)
        }

        u, err := uuid.NewRandom()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, response)
        }

        response.UUID = u.String()
        return c.JSON(http.StatusOK, response)
    })
    e.GET("/regulation", func(c echo.Context) error {
        filename := "./regulation.txt"

        f, err := os.Open(filename)
        defer f.Close()
        
        if err != nil{
            response := regulatoinResponse {
                Message : "File Not Found (" + filename + ")",
            }
            return c.JSON(http.StatusInternalServerError, response)
        }

        b, err := ioutil.ReadAll(f)
        response := regulatoinResponse {
            Message : string(b),
        }
        return c.JSON(http.StatusOK, response)
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
            ID : num,
            Quiz : request.Quiz,
            URL : "https://1.bp.blogspot.com/-3XfMA0UhT70/XlyfrsiokjI/AAAAAAABXn8/j_CLCc73TTEi-PCK19hnUwY3D-pJgmjvQCNcBGAsYHQ/s1600/drink_beer_yukata_man.png",
        }
        return c.JSON(http.StatusOK, response)
    })
    e.Logger.Fatal(e.Start(":1323"))
}
