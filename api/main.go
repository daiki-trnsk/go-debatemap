package main

import (
	"log"
	"os"

	"github.com/daiki-trnsk/go-debatemap/controllers"
	"github.com/daiki-trnsk/go-debatemap/database"
	// customMiddleware "github.com/daiki-trnsk/go-debatemap/middleware"
	"github.com/daiki-trnsk/go-debatemap/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.DebateMapData(database.Client, "DebateMaps"), database.UserData(database.Client, "Users"))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routes.UserRoutes(e)

	e.POST("/signup", controllers.SignUp)
	e.POST("/login", controllers.Login)

	auth := e.Group("")
	auth.GET("/debatemaps_list", app.GetDebateMaps)
	auth.GET("/debatemap/:id", app.GetDebateMap)
	auth.PUT("/debatemap/:id", app.UpdateDebateMap)
	auth.POST("/debatemap", app.AddDebateMap)

	log.Fatal(e.Start(":" + port))
}
