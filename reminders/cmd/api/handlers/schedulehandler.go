package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reminders/cmd/api/vo"
	"reminders/internal/repository"

	"github.com/labstack/echo/v4"
)

func NewSchedule(c echo.Context) error {
	schedule := vo.Schedule{}
	err := json.NewDecoder(c.Request().Body).Decode(&schedule)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your schedule %#v", schedule)
	s := repository.NewScheduleRepository()
	scheduleId := s.NewSchedule(repository.Schedule{
		Id:          repository.GenerateUUID(),
		Description: schedule.Description,
		Users:       schedule.Users,
	})
	log.Printf("The new schedule id is %v", scheduleId)
	return c.String(http.StatusOK, "schedule created successfully!")
}

func UpdateSchedule(c echo.Context) error {
	// id := c.QueryParam("id")
	// description := c.QueryParam("description")
	// users := c.QueryParam("users")
	dataType := c.Param("data")

	if dataType == "json" {
		// schedule := vo.Schedule{
		// 	Id:          id,
		// 	Description: description,
		// 	Users:       users,
		// }
		//TODO: Make the update schedule into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data for the update user",
		})
	}
}

func DeleteSchedule(c echo.Context) error {
	//id := c.QueryParam("id")
	dataType := c.Param("data")

	if dataType == "json" {
		// schedule := vo.Schedule{
		// 	Id:          id,
		// 	Description: "",
		// 	Users:       "",
		// }
		//TODO: Make the delete schedule into data base.
		return c.JSON(http.StatusOK, "")
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetSchedules(c echo.Context) error {
	dataType := c.Param("data")
	if dataType == "json" {
		//schedule := []vo.Schedule; // Call the data base for get all users
		schedule := vo.Schedule{}
		return c.JSON(http.StatusOK, schedule)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Any Error",
		})
	}
}
