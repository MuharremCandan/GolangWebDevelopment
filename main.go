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
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func main() {
	c := echo.New()
	c.Use(setHeader)
	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "satusCode=${status}\n",
	}))

	c.GET("/main", mainHandler)
	c.GET("/user/:data", getUser)

	adminGroup := c.Group("/admin")
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "123" {
			return true, nil
		}
		return false, nil
	}))

	adminGroup.GET("/main", mainAdmin)
	c.POST("/user", addUser)
	c.Start(":8081")
}

func setHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		if contentType == "application/json" {
			return c.String(http.StatusBadRequest, "Yanlızca application/json tipinde istek atılabilir")
		}
		return next(c)
	}
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Admindesin")
}

func mainHandler(c echo.Context) error {

	return c.String(http.StatusOK, "Main handler çağrıldı")
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
