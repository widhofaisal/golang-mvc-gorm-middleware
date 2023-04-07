package controller

import (
	"learn/mvc/model"
	"learn/mvc/config"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []model.User
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})

}

func CreateUser(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}