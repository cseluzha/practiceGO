package main

import (
	"github.com/labstack/echo/v4"
	"reminders/cmd/api/handlers"
)

func main() {
	// create a new echo instance
	e := echo.New()

	//create Routes
	//InitRouting(e)
	// For User
	// Route / handler function
	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.NewUser)
	// For Schedule
	e.GET("/schedules", handlers.GetSchedules)
	e.POST("/schedules", handlers.NewSchedule)
	// //For Output
	// e.GET("/outputs", handlers.GetSchedules)
	// e.POST("/outputs", handlers.NewSchedule)
	e.Logger.Fatal(e.Start(":3000"))
}
