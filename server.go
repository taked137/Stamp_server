package main

import (
    "echo/data"

    "time"
    "os"
    "strconv"
    "fmt"

    "math/rand"
    "net/http"
    "io/ioutil"

    "github.com/labstack/echo/v4"
    "github.com/google/uuid"
)

const (
    BASE_URL string = "http://13.113.250.233:1323"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.GET("/stamp/quiz/:num", func(c echo.Context) error {
        num, _ := strconv.Atoi(c.Param("num"))
        switch num {
            case 1:
                return c.File("./images/quiz1.png")
            case 2:
                return c.File("./images/quiz2.png")
            case 3:
                return c.File("./images/quiz3.png")
            case 4:
                return c.File("./images/quiz4.png")
            case 5:
                return c.File("./images/quiz5.png")
            case 6:
                return c.File("./images/quiz6.png")
            default:
                code := http.StatusNotFound
                msg := http.StatusText(code)
                return c.JSON(code, msg)
        }
    })  
    e.POST("/user/create", func (c echo.Context) error {
        request := new(data.UserRequest)
        if err := c.Bind(request); err != nil {
            return err
        }

        response := data.UserResponse {}

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
            response := data.RegulatoinResponse {
                Message : "File Not Found (" + filename + ")",
            }
            return c.JSON(http.StatusInternalServerError, response)
        }

        b, err := ioutil.ReadAll(f)
        response := data.RegulatoinResponse {
            Message : string(b),
        }
        return c.JSON(http.StatusOK, response)
    })
    e.GET("/stamp/image/:num", func(c echo.Context) error {
        num := c.Param("num")
        response := data.ImageResponse {
            URL : BASE_URL + "/stamp/quiz/" + num,
        }
        return c.JSON(http.StatusOK, response)
    })
    e.POST("/stamp/beacon", func(c echo.Context) error {
        request := new(data.BeaconRequest)
        if err := c.Bind(request); err != nil {
            return err
        }

        rand.Seed(time.Now().UnixNano())
        num := rand.Intn(6)
	    if(1 < num) {
		    num = 2
	    }

        response := data.BeaconResponse {
            ID : num,
            Quiz : request.Quiz,
            //URL : "https://1.bp.blogspot.com/-3XfMA0UhT70/XlyfrsiokjI/AAAAAAABXn8/j_CLCc73TTEi-PCK19hnUwY3D-pJgmjvQCNcBGAsYHQ/s1600/drink_beer_yukata_man.png",
            URL : BASE_URL + "/stamp/quiz/" + strconv.Itoa(request.Quiz),
        }
        return c.JSON(http.StatusOK, response)
    })
    e.POST("/stamp/judge", func(c echo.Context) error {
        request := new(data.AnswerRequest)
        if err := c.Bind(request); err != nil {
            return err
        }

        fmt.Print(c.Request().Header)

        response := data.AnswerResponse {
            Quiz : request.Quiz,
            Correct : false,
        }
        switch(request.Quiz) {
            case 1:
                if(request.Answer == "NIT" || request.Answer == "nit") {
                    response.Correct = true
                }
            case 2:
                if(request.Answer == "c0de") {
                    response.Correct = true
                }
            case 3:
                if(request.Answer == "メイ") {
                    response.Correct = true
                }
            case 4:
                if(request.Answer == "一本松古墳") {
                    response.Correct = true
                }
            case 5:
                if(request.Answer == "57") {
                    response.Correct = true
                }
            case 6:
                if(request.Answer == "はじっこ") {
                    response.Correct = true
                }
        }
        
        return c.JSON(http.StatusOK, response)
    })
    e.POST("/user/goal", func(c echo.Context) error {
        response := data.GoalResponse {
            Accept : true,
        }
        return c.JSON(http.StatusOK, response)
    })
    e.GET("/map/checkpoint", func(c echo.Context) error {
        checkpoints := []data.CheckPoint {
            data.CheckPoint {
                Num : 1,
                Latitude : 35.156893,
		        Longitude : 136.925268,
            },
            data.CheckPoint {
                Num : 2,
                Latitude : 35.158519,
		        Longitude : 136.924672,
            },
            data.CheckPoint {
                Num : 3,
                Latitude : 35.156922,
		        Longitude : 136.926277,
            },
            data.CheckPoint {
                Num : 4,
                Latitude : 35.157514,
		        Longitude : 136.925451,
            },
            data.CheckPoint {
                Num : 5,
                Latitude : 35.156142,
		        Longitude : 136.924496,
            },
            data.CheckPoint {
                Num : 6,
                Latitude : 35.157689,
		        Longitude : 136.924184,
            },
        }
        response := data.MapResponse {
            Point : checkpoints,
        }

        return c.JSON(http.StatusOK, response)
    })
    e.GET("/info", func(c echo.Context) error {
        offset, _ := strconv.Atoi(c.QueryParam("offset"))
        if offset < 0 {
            return nil
        }
        
        limit, _ := strconv.Atoi(c.QueryParam("limit"))
        if limit == 0 {
            limit = 300
        }

        messages := make([]data.InfoResponse, limit)

        categories := []string{"お化け屋敷", "出店", "研究室見学", "図書館"}

        for i := 0; i < limit; i++ {
            rand.Seed(time.Now().UnixNano())
            num := rand.Intn(len(categories))

            messages[i] = data.InfoResponse {
                Message : ("[" + categories[num] + "] " + strconv.Itoa(i + offset)),
            }
        } 

        response := make(map[string][]data.InfoResponse)
        response["result"] = messages
        return c.JSON(http.StatusOK, response)
    })

    e.Logger.Fatal(e.Start(":1323"))
}
