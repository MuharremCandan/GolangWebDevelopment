package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Username string "json:'username'"
	Name     string "json:'name'"
	Surname  string "json:'surname'"
}

func main() {
	c := echo.New()
	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "satusCode=${status}\n",
	}))

	c.GET("/main", mainHandler)
	c.GET("/user/:data", getUser)

	adminGroup := c.Group("/admin")

	adminGroup.GET("/main", mainAdmin)
	c.POST("/user", addUser)
	c.Start(":8081")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Admindesin")
}

func mainHandler(c echo.Context) error {

	return c.String(http.StatusOK, "İlk backend project")
}

func getUser(c echo.Context) error {

	dataType := c.Param("data")

	userName := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("User Name: %s \nName: %s\nSurname: %s", userName, name, surname))
	}

	if dataType == "json" {

		return c.JSON(http.StatusOK, map[string]string{
			"username": userName,
			"name":     name,
			"surname:": surname,
		})

	}
	return c.String(http.StatusBadRequest, "JSON ve String Formatlarında Dönüş Yapıınabilinir !")
}

func addUser(c echo.Context) error {
	user := new(User)
	body, error := ioutil.ReadAll(c.Request().Body)

	if error != nil {
		return error
	}

	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	fmt.Println(user)
	return c.String(200, "Yaşıyoo")
}
