package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reminders/cmd/api/vo"
	"reminders/internal/repository"

	//"reminders/internal/repository"

	"github.com/labstack/echo/v4"
)

func NewUser(c echo.Context) error {
	user := vo.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	defer c.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your user %v", user)
	//TODO: call the create new user into data base
	g := repository.NewUserRepository()
	userid := g.NewUser(repository.User{
		IdUser: repository.GenerateUUID(),
		Email:  user.Email,
	})
	log.Printf("The new user id is %v", userid)
	return c.String(http.StatusOK, "We got your User: %v")
}

func UpdateUser(c echo.Context) error {
	// userId := c.QueryParam("userid")
	// email := c.QueryParam("email")
	// dataType := c.Param("data")

	user := vo.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your user %#v to update", user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	userId := c.QueryParam("userid")
	dataType := c.Param("data")

	if dataType == "json" {
		user := vo.User{
			IdUser: userId,
			Email:  "",
		}
		//TODO: Make the delete user into data base.
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data",
		})
	}
}

func GetUsers(c echo.Context) error {
	dataType := c.Param("data")
	if dataType == "json" {
		//users := []vo.User; // Call the data base for get all users
		user := vo.User{}
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Any Error",
		})
	}
}
