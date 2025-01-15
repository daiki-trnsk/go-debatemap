package main

import (
	"log"
	"os"

	"github.com/daiki-trnsk/go-debatemap/api/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signup", controllers.SignUp)
	e.POST("/login", controllers.Login)

	auth := e.Group("")
	auth.GET("/debatemaps_list", controllers.GetDebateMaps)
	auth.GET("/debatemap/:id", controllers.GetDebateMap)
	auth.PUT("/update/:id", controllers.UpdateDebateMap)
	auth.POST("/adddebatemap", controllers.AddDebateMap)

	log.Fatal(e.Start(":" + port))
}
