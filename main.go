package main

import (
  "net/http"
	"github.com/labstack/echo/v4"
)

func main(){
  e := echo.New()

  e.GET("/index", func(c echo.Context) error{
    return c.String(200, "ini index")
  })

  e.POST("/index", func(c echo.Context) error{
    return c.String(400, "ini index")
  })

  e.PUT("/index", func(c echo.Context) error{
    return c.String(400, "ini index")
  })

  e.DELETE("/index", func(c echo.Context) error{
    return c.String(400, "ini index")
  })
  
  e.GET("/index/:id", func(c echo.Context) error{
    id := c.Param("id")
    return c.String(http.StatusPaymentRequired, "Nama " + id)
  })

  e.GET("/index/:id/:id2", func(c echo.Context) error{
    id := c.Param("id")
    id2 := c.Param("id2")
    return c.String(http.StatusPaymentRequired, "Nama " + id + " NIM " + id2)
  })

  e.Start(":2121")
}